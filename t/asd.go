package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"io/ioutil"
	"log"
	"strconv"
	"time"
)

func main()  {
	f, _ :=ioutil.ReadFile("C:\\Users\\Beaver\\Pictures\\1.png")

	var strRequestURI = []byte("http://localhost:8080/template/upload")

	req := fasthttp.AcquireRequest()

	req.AppendBody(f)
	req.AppendBody(f)
	req.AppendBody(f)

	var strPost = []byte("POST")
	req.Header.SetMethodBytes(strPost)
	req.SetRequestURIBytes(strRequestURI)
	res := fasthttp.AcquireResponse()
	if err := fasthttp.Do(req, res); err != nil {
		panic("handle error")
	}
	fasthttp.ReleaseRequest(req)

	fmt.Println(string(res.Body()))
}


func TestUpload2() {
	time.Sleep(5000 * time.Millisecond)
	testFile1, _ := ioutil.ReadFile("./test/test1.png")
	log.Println(len(testFile1))
	testFile2, _ := ioutil.ReadFile("./test/test2.png")
	log.Println(len(testFile2))
	testFile3, _ := ioutil.ReadFile("./test/test3.png")
	log.Println(len(testFile3))

	req := *fasthttp.AcquireRequest()
	res := *fasthttp.AcquireResponse()

	req.SetRequestURI("http://localhost:8080/template/upload")
	req.Header.SetMethod("POST")
	req.Header.Add("file_1.png", strconv.Itoa(len(testFile1)))
	req.Header.Add("file_2.png", strconv.Itoa(len(testFile2)))
	req.Header.Add("file_3.png", strconv.Itoa(len(testFile3)))

	req.AppendBody(testFile1)
	req.AppendBody(testFile2)
	req.AppendBody(testFile3)

	if err := fasthttp.Do(&req, &res); err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(res.Body()))

	matchReq := *fasthttp.AcquireRequest()
	matchRes := *fasthttp.AcquireResponse()

	url := "http://localhost:8080/template/detect/" + string(res.Body())
	matchReq.SetRequestURI(url)
	matchReq.Header.SetMethod("POST")
	matchReq.SetBody(testFile3)

	if err := fasthttp.Do(&matchReq, &matchRes); err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(matchRes.Body()))

}