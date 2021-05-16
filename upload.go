package main

import (
	"bytes"
	"fmt"
	"github.com/valyala/fasthttp"
	"gocv.io/x/gocv"
	"io"
)

func uploadTemplatesHandler(ctx *fasthttp.RequestCtx) {
	form, err := ctx.Request.MultipartForm()
	if err != nil {
		ctx.Error(err.Error(), 400)
		return
	}
	var templates []gocv.Mat
	for _, multipart := range form.File {
		for _, file := range multipart {
			f, _ := file.Open()
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

	if len(templates) == 0 {
		ctx.Error("Can't load templates", 400)
		return
	}

	templatesId := AddTemplates(templates)

	response := fmt.Sprintf("{\"templateId\":\"%s\"}", templatesId)
	ctx.Response.SetBodyString(response)
	ctx.SetContentType("application/json")
}
