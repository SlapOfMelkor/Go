package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"
	"sync"
)

const (
	liste  = "liste.txt"
	worker = 100
)

var hedefurl string

// test icin kullandigim site "https://rust.facepunch.com/"----------------------------------------------
func init() {
	flag.StringVar(&hedefurl, "u", "", "Hedef URL (ornegin ,go run dosyaadi.go -u <hedef_url>)")
	flag.Parse()

	if hedefurl == "" {
		fmt.Println("Kullanim: 'go run dosyaadi.go -u <hedef_url>'")
		os.Exit(1)
	}
}

func main() {
	var wg sync.WaitGroup
	file, _ := os.Open(liste)

	defer file.Close()

	var wordlist []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		wordlist = append(wordlist, word)
	}

	jobs := make(chan string, 100)
	results := make(chan string, 100)

	for i := 0; i < worker; i++ {
		go func() {
			for word := range jobs {
				url := fmt.Sprintf("%s/%s", hedefurl, word)
				resp, err := http.Get(url)
				if err != nil {
					fmt.Println("Istek hatasi", err)
					wg.Done()
					continue
				}
				defer resp.Body.Close()
				statusCode := resp.StatusCode
				if statusCode == 200 {
					result := fmt.Sprintf("%s", url)
					results <- result
				}
				wg.Done()
			}
		}()
	}

	for _, word := range wordlist {
		wg.Add(1)
		jobs <- word
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	fmt.Println("Wordlistteki girdilere gore girdiginiz websitesinde mevcut sayfalar sunlar; -----------------------------")
	for result := range results {
		fmt.Println(result)
	}
}
