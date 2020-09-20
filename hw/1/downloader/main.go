// 2. * Напишите функцию, которая получает на вход публичную ссылку на файл с «Яндекс.Диска
// и сохраняет полученный файл на диск пользователя.

// Разобраться с авторизациеи и получением токена

package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	// linkPath := downloadApi + "?" + downloadPath

	if err := getReq(link); err != nil {
		log.Fatal(err)
	}
}

const (
	downloadApi  = "https://cloud-api.yandex.net/v1/disk/resources/download"
	downloadPath = "https://yadi.sk/i/uMSI3I6C6VdmhA"
	link         = "https://downloader.dst.yandex.ru/disk/53139aa0et584d3bac7eeab405d3574b/535320b4/YyjTJtEHob8R5WbpojJbiiUuU2HC_2JSTU0gW9qE0NHGW2uncmBjM_-IXun3Msyij96FTHQGSX-fDL-XwokDvA%3D%3D?uid=202727674&filename=photo.png&disposition=attachment&hash=&limit=0&content_type=application%2Fx-www-form-urlencoded&fsize=34524&hid=93528043563b8r55723a253f4730290a&media_type=document"
)

func getReq(Llink string) error {
	resp, err := http.Get(Llink)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	log.Printf("Status: %v; \nStatusCode: %v; \nHeader: %v\n", resp.Status, resp.StatusCode, resp.Header)
	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		return err
	}
	return nil
}
