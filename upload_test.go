package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"io/ioutil"
	"log"
	"strconv"
	"testing"
)

func TestShouldParseImages(t *testing.T)  {
	testFile1, _ :=ioutil.ReadFile("./test/test1.png")
	log.Println(len(testFile1))
	testFile2, _ :=ioutil.ReadFile("./test/test2.png")
	log.Println(len(testFile2))
	testFile3, _ :=ioutil.ReadFile("./test/test3.png")
	log.Println(len(testFile3))

	var body = append(testFile1, testFile2...)
	body = append(body, testFile3...)

	req := fasthttp.AcquireRequest()

	req.Header.Add("test1.png", strconv.Itoa(len(testFile1)))
	req.Header.Add("test2.png", strconv.Itoa(len(testFile2)))
	req.Header.Add("test3.png", strconv.Itoa(len(testFile3)))

	var parsedImages = parseImages(&req.Header,body)

	if 3 != len(parsedImages) {
		t.Errorf("images after parsing are not the same.")
	}
}

func TestUpload(t *testing.T)  {
	testFile1, _ :=ioutil.ReadFile("./test/test1.png")
	log.Println(len(testFile1))
	testFile2, _ :=ioutil.ReadFile("./test/test2.png")
	log.Println(len(testFile2))
	testFile3, _ :=ioutil.ReadFile("./test/test3.png")
	log.Println(len(testFile3))


	req := *fasthttp.AcquireRequest()
	res := *fasthttp.AcquireResponse()

	req.SetRequestURI("http://localhost:8080/template/upload")
	req.Header.SetMethod("POST")
	req.Header.Add("test1.png", strconv.Itoa(len(testFile1)))
	req.Header.Add("test2.png", strconv.Itoa(len(testFile2)))
	req.Header.Add("test3.png", strconv.Itoa(len(testFile3)))

	req.AppendBody(testFile1)
	req.AppendBody(testFile2)
	req.AppendBody(testFile3)

	if err := fasthttp.Do(&req, &res); err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(res.Body()))
	fmt.Println(string(res.Body()))
}