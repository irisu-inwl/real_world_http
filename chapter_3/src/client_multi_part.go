package main

import (
	"bytes"
	"io"
	"mime/multipart"
	"log"
	"os"
	"net/http"
	"fmt"
	"net/textproto"
)

func main() {
	host := os.Getenv("HOST")
	url := fmt.Sprintf("http://%s:18888", host)

	// Setting multi-part body
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	writer.WriteField("name", "Michael Jackson")

	// Setting multi-part header and file to send
	part := make(textproto.MIMEHeader)
	part.Set("Content-Type", "image/jpeg")
	part.Set("Content-Disposition", `form-data; name="thumbnail"; filename="photo.png"`)
	fileWriter, err := writer.CreatePart(part)
	if err != nil {
		panic(err)
	}
	readFile, err := os.Open("photo.png")
	if err != nil {
		panic(err)
	}
	io.Copy(fileWriter, readFile)
	writer.Close()

	resp, err := http.Post(url, writer.FormDataContentType(), &buffer)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	// log.Println(string(body))
	log.Println("Status:", resp.Status)
}
