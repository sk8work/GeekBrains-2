// Request Line - Get http://ya.ru/ HTTP/1.1

// Accept: text/html, text/plan, application/json
// Accept-language: ru-RU, ru; q=0.9, en-US; q=0.8
// Connection: keep-alive( или close)
// Host: ya.ru
// User-Agent: Mozilla/5.0 (.... - операционные системы)
//

package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	if err := getReq(); err != nil {
		log.Fatal(err)
	}
}

const (
	getURL = "http://ya.ru"
)

func getReq() error {
	resp, err := http.Get(getURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	log.Printf("Status: %v; \nStatusCode: %v; \nHeader: %v\n", resp.Status, resp.StatusCode, resp.Header)

	// _, err := io.Copy(os.Stdout, resp.Body)
	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		return err
	}
	return nil
}
