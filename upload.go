package main

import (
	"bytes"
	"fmt"
	"github.com/google/uuid"
	"github.com/valyala/fasthttp"
	"gocv.io/x/gocv"
	"io"
	"log"
	"strconv"
	"strings"
)

func uploadTemplatesHandler(ctx *fasthttp.RequestCtx) {
	form, err := ctx.Request.MultipartForm()
	if err != nil {
		ctx.Error(err.Error(), 400)
		return
	}
	var templates []gocv.Mat
	for _, v := range form.File {
		for _, t := range v {
			f, _ := t.Open()
			b := new(bytes.Buffer)
			_, er := io.Copy(b, f)
			if er != nil {
				ctx.Error(er.Error(), 400)
				return
			}
			if template, err := gocv.IMDecode(b.Bytes(), gocv.IMReadGrayScale); err == nil {
				templates = append(templates, template)
			}

		}
	}
	// templates := parseImages(&ctx.Request.Header, body)
	if len(templates) == 0 {
		ctx.Error("Can't load templates", 400)
		return
	}
	templateId := uuid.New().String()
	templatesMap[templateId] = templates

	response := fmt.Sprintf("{\"templateId\":\"%s\"}", templateId)

	ctx.Response.SetBodyString(response)

	ctx.SetContentType("application/json")

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
