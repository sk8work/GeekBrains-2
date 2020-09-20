// Request Line - POST http://ya.ru/ HTTP/1.1

// Accept: text/html, text/plan, application/json
// Accept-language: ru-RU, ru; q=0.9, en-US; q=0.8
// Connection: keep-alive( или close)
// Host: ya.ru
// User-Agent: Mozilla/5.0 (.... - операционные системы)
// Content-Type: application/json

// 1) Request Line
// 2) Headers
// 3) Body

package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	if err := postReq(); err != nil {
		log.Fatal(err)
	}
}

const (
	getURL  = "http://ya.ru"
	postURL = "https://postman-echo.com/post"
)

func postReq() error {
	resp, err := http.Post(
		postURL, "application/json", strings.NewReader(`{"foo1":"bar1", "foo2":"bar2"}`))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		return err
	}
	return nil
}
