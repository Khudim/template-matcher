package main

import (
	"bytes"
	"errors"
	"github.com/valyala/fasthttp"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"
	"testing"
)

func TestShouldParseImages(t *testing.T) {
	testFile1, _ := ioutil.ReadFile("./test/test1.png")
	log.Println(len(testFile1))
	testFile2, _ := ioutil.ReadFile("./test/test2.png")
	log.Println(len(testFile2))
	testFile3, _ := ioutil.ReadFile("./test/test3.png")
	log.Println(len(testFile3))

	var body = append(testFile1, testFile2...)
	body = append(body, testFile3...)

	req := fasthttp.AcquireRequest()

	/*	req.Header.Add("test1.png", strconv.Itoa(len(testFile1)))
		req.Header.Add("test2.png", strconv.Itoa(len(testFile2)))
		req.Header.Add("test3.png", strconv.Itoa(len(testFile3)))
	*/
	var parsedImages = parseImages(&req.Header, body)

	if 3 != len(parsedImages) {
		t.Errorf("images after parsing are not the same.")
	}
}

func TestUpload(t *testing.T) {
	testFile1, _ := ioutil.ReadFile("./test/test1.png")
	log.Println(len(testFile1))
	testFile2, _ := ioutil.ReadFile("./test/test2.png")
	log.Println(len(testFile2))
	testFile3, _ := ioutil.ReadFile("./test/test3.png")
	log.Println(len(testFile3))

	/*	req := *fasthttp.AcquireRequest()
		res := *fasthttp.AcquireResponse()

		req.SetRequestURI("http://localhost:8080/template/upload")
		req.Header.SetMethod("POST")*/
	/*
		var body = append(testFile1, testFile2...)
		body = append(body, testFile3...)
	*/
	err := UploadFile("http://localhost:8080/template/upload", []string{"./test/test1.png", "./test/test2.png", "./test/test3.png"})
	if err != nil {
		log.Println(err)
	}
	/*form, err := mr.ReadForm(1024)
	form2, err := mr2.ReadForm(1024)
	form3, err := mr3.ReadForm(1024)

	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	*/
	/*	if err := fasthttp.WriteMultipartForm(&w, form, "f1"); err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		if err := fasthttp.WriteMultipartForm(&w, form2, "f2"); err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		if err := fasthttp.WriteMultipartForm(&w, form3, "f3"); err != nil {
			t.Fatalf("unexpected error: %s", err)
		}*/
	/*	req.AppendBody(testFile1)
		req.AppendBody(testFile2)
		req.AppendBody(testFile3)*/

	/*	if err := fasthttp.Do(&req, &res); err != nil {
			fmt.Println(err)
		}

		fmt.Println(string(res.Body()))
		fmt.Println(string(res.Body()))*/
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
