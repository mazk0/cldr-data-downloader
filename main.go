package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const outputFolderName = "cldr-data"

func main() {
	body := download("https://github.com/unicode-cldr/cldr-core/archive/36.0.0.zip")
	deflate(body)
	body = download("https://github.com/unicode-cldr/cldr-dates-modern/archive/36.0.0.zip")
	deflate(body)
	body = download("https://github.com/unicode-cldr/cldr-cal-buddhist-modern/archive/36.0.0.zip")
	deflate(body)
	body = download("https://github.com/unicode-cldr/cldr-cal-chinese-modern/archive/36.0.0.zip")
	deflate(body)
	body = download("https://github.com/unicode-cldr/cldr-cal-coptic-modern/archive/36.0.0.zip")
	deflate(body)
	body = download("https://github.com/unicode-cldr/cldr-cal-dangi-modern/archive/36.0.0.zip")
	deflate(body)
	body = download("https://github.com/unicode-cldr/cldr-cal-ethiopic-modern/archive/36.0.0.zip")
	deflate(body)
	body = download("https://github.com/unicode-cldr/cldr-cal-hebrew-modern/archive/36.0.0.zip")
	deflate(body)
	body = download("https://github.com/unicode-cldr/cldr-cal-indian-modern/archive/36.0.0.zip")
	deflate(body)
	body = download("https://github.com/unicode-cldr/cldr-cal-islamic-modern/archive/36.0.0.zip")
	deflate(body)
	body = download("https://github.com/unicode-cldr/cldr-cal-japanese-modern/archive/36.0.0.zip")
	deflate(body)
	body = download("https://github.com/unicode-cldr/cldr-cal-persian-modern/archive/36.0.0.zip")
	deflate(body)
	body = download("https://github.com/unicode-cldr/cldr-cal-roc-modern/archive/36.0.0.zip")
	deflate(body)
	body = download("https://github.com/unicode-cldr/cldr-localenames-modern/archive/36.0.0.zip")
	deflate(body)
	body = download("https://github.com/unicode-cldr/cldr-misc-modern/archive/36.0.0.zip")
	deflate(body)
	body = download("https://github.com/unicode-cldr/cldr-numbers-modern/archive/36.0.0.zip")
	deflate(body)
	body = download("https://github.com/unicode-cldr/cldr-segments-modern/archive/36.0.0.zip")
	deflate(body)
	body = download("https://github.com/unicode-cldr/cldr-units-modern/archive/36.0.0.zip")
	deflate(body)
}

func download(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	return body
}

func deflate(body []byte) {
	reader, err := zip.NewReader(bytes.NewReader(body), int64(len(body)))
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range reader.File {
		if strings.Contains(file.Name, "..") {
			fmt.Printf("File %v contained illegal path operation .. processing next file.", file.Name)
			continue
		} else {
			writeFileToDisk(file)
		}
	}
}

func writeFileToDisk(file *zip.File) {
	indexOfFirstSlash := strings.Index(file.Name, "/")
	filePath := outputFolderName + "/" + file.Name[indexOfFirstSlash:]
	if file.FileInfo().IsDir() {
		err := os.Mkdir(filePath, 0755)
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	outFile, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err)
	}
	defer outFile.Close()

	zipFile, err := file.Open()
	if err != nil {
		fmt.Println(err)
	}
	defer zipFile.Close()

	io.Copy(outFile, zipFile)
}
