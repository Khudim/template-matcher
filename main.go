package main

import (
	"github.com/fasthttp/router"
	"log"
	"os"
)
import "github.com/valyala/fasthttp"

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
