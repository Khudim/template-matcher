package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"gocv.io/x/gocv"
	"image"
	"log"
)

func detectHandler(ctx *fasthttp.RequestCtx) {
	templateId := ctx.UserValue("templateId")

	templates := templatesMap[templateId.(string)]
	result := make(chan gocv.Mat, len(templates))

	if img, err := gocv.IMDecode(ctx.Request.Body(), gocv.IMReadGrayScale); err == nil {

		for _, template := range templates {
			go findTemplate(template, img, result)
		}
	} else {
		fmt.Println(err)
		ctx.Response.SetStatusCode(400)
		return
	}
	if templateId != nil {
		fmt.Println(templateId)
	}

	maxConf, bestPoint := float32(0), image.Point{}

	for i := 0; i < len(templates); i++ {
		select {
		case point := <-result:
			_, maxConfidence, _, p := gocv.MinMaxLoc(point)
			if maxConf < maxConfidence {
				maxConf, bestPoint = maxConfidence, p
			}
			closeResource(point)
		}
	}
	response := fmt.Sprintf("{\"confidence\":%f,\"x\":%d,\"y\":%d}", maxConf, bestPoint.X, bestPoint.Y)
	ctx.Response.AppendBodyString(response)
	ctx.Response.SetStatusCode(200)
}

func findTemplate(template, img gocv.Mat, resultChannel chan gocv.Mat) {
	result := gocv.NewMat()

	mask := gocv.NewMat()
	defer closeResource(mask)

	gocv.MatchTemplate(img, template, &result, 5, mask)
	resultChannel <- result
}

func closeResource(resource gocv.Mat) {
	if err := resource.Close(); err != nil {
		log.Println(err)
	}
}
