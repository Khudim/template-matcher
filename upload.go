package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/valyala/fasthttp"
	"gocv.io/x/gocv"
	"log"
	"strconv"
	"strings"
)

func uploadTemplatesHandler(ctx *fasthttp.RequestCtx) {
	body := ctx.Request.Body()
	templates := parseImages(&ctx.Request.Header, body)
	if len(templates) == 0 {
		ctx.Response.SetStatusCode(400)
		return
	}
	templatesId := uuid.New().String()
	templatesMap[templatesId] = templates

	ctx.Response.SetBody([]byte(templatesId))

	ctx.SetContentType("application/octet-stream")

}

func parseImages(headers *fasthttp.RequestHeader, body []byte) []gocv.Mat {
	offset := 0
	var templates []gocv.Mat
	headers.VisitAll(func(name, size []byte) {
		fileName := strings.ToLower(string(name))
		if strings.HasPrefix(fileName, "file_") {
			limit, _ := strconv.Atoi(string(size))
			if offset+limit > len(body) {
				log.Fatalln(fmt.Errorf("image size more than request body, name = %s, size = %d", name, limit))
			}
			image := body[offset : offset+limit]
			if template, err := gocv.IMDecode(image, gocv.IMReadGrayScale); err == nil {
				templates = append(templates, template)
			}
			offset += limit
		}
	})
	return templates
}
