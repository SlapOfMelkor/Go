package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Kullanım: go run main.go <url> <aramaMetni>")
		os.Exit(1)
	}

	url := os.Args[1]
	aramaMetni := os.Args[2]

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("URL alınamıyor: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Hata: %s adresi %d status koduyla döndü\n", url, resp.StatusCode)
		os.Exit(1)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("HTML okunamıyor: %v\n", err)
		os.Exit(1)
	}

	index := strings.Index(string(body), aramaMetni)
	if index != -1 && index+len(aramaMetni)+6 < len(body) {
		substring := string(body[index+len(aramaMetni) : index+len(aramaMetni)+6])
		var numericString string
		for _, char := range substring {
			if unicode.IsDigit(char) {
				numericString += string(char)
			}
		}
		lastint, _ := strconv.Atoi(numericString)

		fmt.Printf("Girilen Sitenin Wordpress Surumu= %d\n", lastint)
	} else {
		fmt.Printf("Bu sitenin Wordpress Surumu Bulunamadi")
	}
}
