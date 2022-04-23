package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	host := os.Getenv("HOST")
	url := fmt.Sprintf("http://%s:18888/chunked", host)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	reader := bufio.NewReader(resp.Body)
	for {
		line, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		log.Println(string(bytes.TrimSpace(line)))
	}
}
