package main

import (
	"io/ioutil"
	"log"
	"os"
	"net/http"
	"fmt"
)

func main() {
	host := os.Getenv("HOST")
	url := fmt.Sprintf("http://%s:18888", host)
	resp, err := http.Get(url)
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
