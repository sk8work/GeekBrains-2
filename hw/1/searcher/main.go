// 1. Напишите функцию, которая будет получать на вход строку с поисковым запросом (string) и
// массив ссылок на страницы, по которым стоит произвести поиск ([]string). Результатом работы
// функции должен быть массив строк со ссылками на страницы, на которых обнаружен
// поисковый запрос. Функция должна искать точное соответствие фразе в тексте ответа от сервера по каждой из ссылок.
// Подсказка: это задача на поиск последовательности в массиве

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

const (
	textReq = "Погода"
)

func main() {

	urls := make([]string, 0, 10)
	urls = append(urls, "http://yandex.ru")
	urls = append(urls, "http://ya.ru")
	urls = append(urls, "http://beeline.ru")
	urls = append(urls, "http://yahoo.com")
	urls = append(urls, "http://mail.ru")
	urls = append(urls, "http://google.ru")
	urls = append(urls, "http://geekbrains.ru")

	for i := 0; i < len(urls); i++ {
		if err := getReq(urls[i], textReq); err != nil {
			log.Fatal(err)
		}
	}
}

func getReq(getURL, textReq string) error {
	resp, err := http.Get(getURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	htmlData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(string(htmlData))
	verified, err := regexp.MatchString(textReq, string(htmlData))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("is verified url='%v' \tby '%v': %v\n", getURL, textReq, verified)

	return nil
}
