package main

import (
	"io/ioutil"
	"log"
	"os"
	"net/http"
	"fmt"
	"strings"
)

func main() {
	host := os.Getenv("HOST")
	reader := strings.NewReader("text")

	url := fmt.Sprintf("http://%s:18888", host)

	resp, err := http.Post(url, "text/plain", reader)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
}
