package main

import (
	"bytes"
	"errors"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"
	"testing"
)

func TestUpload(t *testing.T) {
	testFile1, _ := ioutil.ReadFile("./test/test1.png")
	log.Println(len(testFile1))
	testFile2, _ := ioutil.ReadFile("./test/test2.png")
	log.Println(len(testFile2))
	testFile3, _ := ioutil.ReadFile("./test/test3.png")
	log.Println(len(testFile3))

	err := UploadFile("http://localhost:8080/template/upload", []string{"./test/test1.png", "./test/test2.png", "./test/test3.png"})
	if err != nil {
		log.Println(err)
	}
}

func UploadFile(uri string, files []string) error {

	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)

	for i, f := range files {
		part, err := writer.CreateFormFile("f"+strconv.Itoa(i), "./test/test"+strconv.Itoa(i)+".png")
		if err != nil {
			log.Println(err)
			return err
		}
		b, err := ioutil.ReadFile(f)
		if err != nil {
			log.Println(err)
			return err
		}
		part.Write(b)
	}
	writer.Close()

	req, _ := http.NewRequest("POST", uri, buf)

	req.Header.Add("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	r, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode >= 400 {
		return errors.New(string(r))
	}
	log.Println(string(r))
	return nil
}
