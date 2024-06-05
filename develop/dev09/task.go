package main

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/pborman/getopt"
)

// getURLandFilename извлекает URL и имя файла из командной строки, парсирует URL и возвращает его и имя файла.
func getURLandFilename() (string, string) {
	urlPath := getopt.StringLong("url", 'u', "", "url")
	getopt.Parse()
	_, err := url.Parse(*urlPath)
	if err!= nil {
		log.Fatal(err)
	}
	fmt.Println(*urlPath)
	splitedURL := strings.Split(*urlPath, "/")
	return *urlPath, splitedURL[len(splitedURL)-1]
}

// createFile создает новый файл с указанным именем и возвращает дескриптор этого файла.
func createFile(filename string) *os.File {
	fmt.Printf("filename: %s\n", filename)
	file, err := os.Create(filename)
	if err!= nil {
		log.Fatal(err)
	}
	return file
}

// getData скачивает данные с указанного URL и записывает их в файл, возвращая размер скачанных данных.
func getData(urlPath string, client *http.Client, file *os.File) int64 {
	resp, err := client.Get(urlPath)
	if err!= nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	size, err := io.Copy(file, resp.Body)
	if err!= nil {
		log.Fatal(err)
	}
	defer file.Close()
	return size
}

func main() {
	// Получение URL и имени файла из командной строки.
	urlPath, filename := getURLandFilename()
	fmt.Println("Loh")
	// Создание файла для записи данных.
	file := createFile(filename)
	fmt.Println("Hol")
	// Настройка клиента HTTP для обработки перенаправлений.
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	// Скачивание данных с URL и запись их в файл.
	size := getData(urlPath, &client, file)

	fmt.Printf("Downloaded a file %s with size %d", urlPath, size)
}

