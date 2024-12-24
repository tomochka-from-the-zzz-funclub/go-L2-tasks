package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func GetPage(url, filename string) error {
	resp, err1 := http.Get(url)
	if err1 != nil {
		return err1
	}
	defer resp.Body.Close()
	f, err2 := os.Create(filename)
	if err2 != nil {
		return err2
	}
	defer f.Close()
	_, err3 := io.Copy(f, resp.Body)
	if err3 != nil {
		return err2
	}
	return nil
}

func main() {
	fmt.Print("Please input url of website: ") //  https://mai.ru/
	var url string
	fmt.Fscan(os.Stdin, &url)
	filename := "web.txt"
	err := GetPage(url, filename)
	if err != nil {
		log.Fatal(err)
	}
}
