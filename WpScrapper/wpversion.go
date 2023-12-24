package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
)

type wpApi struct {
	Error   int         `json:"error"`
	Message interface{} `json:"message"`
	Data    struct {
		Core          string      `json:"core"`
		Link          interface{} `json:"link"`
		Vulnerability []struct {
			Name        string      `json:"name"`
			Description interface{} `json:"description"`
			Source      []struct {
				ID          string `json:"id"`
				Name        string `json:"name"`
				Link        string `json:"link"`
				Description string `json:"description"`
				Date        string `json:"date"`
			} `json:"source"`
			Impact struct {
				Cvss struct {
					Version     string `json:"version"`
					Vector      string `json:"vector"`
					Av          string `json:"av"`
					Ac          string `json:"ac"`
					Pr          string `json:"pr"`
					UI          string `json:"ui"`
					S           string `json:"s"`
					C           string `json:"c"`
					I           string `json:"i"`
					A           string `json:"a"`
					Score       string `json:"score"`
					Severity    string `json:"severity"`
					Exploitable string `json:"exploitable"`
					Impact      string `json:"impact"`
				} `json:"cvss"`
				Cwe []struct {
					Cwe         string `json:"cwe"`
					Name        string `json:"name"`
					Description string `json:"description"`
				} `json:"cwe"`
			} `json:"impact"`
		} `json:"vulnerability"`
	} `json:"data"`
	Updated string `json:"updated"`
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Kullaim: go run main.go <url> Wordpress")
		os.Exit(1)
	}

	url := os.Args[1]
	aramaMetni := os.Args[2]

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("URL alinamiyor: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Hata: %s adresi %d status koduyla döndü\n", url, resp.StatusCode)
		os.Exit(1)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("HTML okunamiyor: %v\n", err)
		os.Exit(1)
	}
	var substring string
	index := strings.Index(string(body), aramaMetni)
	if index != -1 && index+len(aramaMetni)+6 < len(body) {
		substring = string(body[index+len(aramaMetni) : index+len(aramaMetni)+6])

		fmt.Println(substring)
	} else {
		fmt.Printf("Bu sitenin Wordpress Surumu Bulunamadi")
	}
	apiURL := "https://www.wpvulnerability.net/core/" + "substring" // eger ki bu kod ile calisan bir websitesi olmazsa test icin substring yerine "5.2.1" yazilarak geri kalan kisim test edilebilir

	apiResp, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("API'ye bağlanirken bir hata oluştu:", err)
		return
	}
	defer apiResp.Body.Close()
	apiBody, err := io.ReadAll(apiResp.Body)
	if err != nil {
		fmt.Println("API yanitini okurken bir hata oluştu:", err)
		return
	}
	var api wpApi
	json.Unmarshal(apiBody, &api)
	printTable(api)
}
func printTable(data wpApi) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Vulnerability Name", "CVE ID", "Description", "Source", "Source ID", "CVSS Score", "Severity", "CWE"})
	for _, vuln := range data.Data.Vulnerability {
		for _, source := range vuln.Source {
			var cveID string
			if len(source.ID) > 0 {
				cveID = source.ID
			} else {
				cveID = "N/A"
			}
			var description string
			if vuln.Description != nil {
				description = fmt.Sprintf("%v", vuln.Description)
			} else {
				description = "N/A"
			}
			var sourceName string
			if len(source.Name) > 0 {
				sourceName = source.Name
			} else {
				sourceName = "N/A"
			}
			cvssScore := vuln.Impact.Cvss.Score
			severity := vuln.Impact.Cvss.Severity
			var cwe string
			if len(vuln.Impact.Cwe) > 0 {
				cwe = vuln.Impact.Cwe[0].Cwe
			} else {
				cwe = "N/A"
			}
			table.Append([]string{vuln.Name, cveID, description, sourceName, source.ID, cvssScore, severity, cwe})
		}
	}
	table.Render() //terminalde tam ekran seklinde en okunakli hali bastirilir!!!!
}
