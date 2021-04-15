package main

import (
	"github.com/fasthttp/router"
	"gocv.io/x/gocv"
	"log"
	"os"
)
import "github.com/valyala/fasthttp"

var templatesMap = make(map[string][]gocv.Mat)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Started")

	r := router.New()
	r.GET("/ping", func(ctx *fasthttp.RequestCtx) {
		ctx.Response.SetBodyString("Alive")
	})
	r.POST("/template/upload", uploadTemplatesHandler)
	r.POST("/template/detect/{templateId}", detectHandler)

	log.Fatal(fasthttp.ListenAndServe(":"+port, r.Handler))
}
