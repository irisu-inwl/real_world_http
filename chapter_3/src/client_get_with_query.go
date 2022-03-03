package main

import (
	"io/ioutil"
	"log"
	"os"
	"net/http"
	"fmt"
	"net/url"
)

func main() {
	host := os.Getenv("HOST")
	values := url.Values{
		"query": {"hello world"},
	}

	url := fmt.Sprintf("http://%s:18888?%s", host, values.Encode())
		
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
