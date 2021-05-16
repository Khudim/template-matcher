package main

import (
	"github.com/google/uuid"
	"gocv.io/x/gocv"
	"time"
)

var templatesMap = make(map[string]*templatesInfo)

type templatesInfo struct {
	templates []gocv.Mat
	usedAt    time.Time
}

func GetTemplates(templateId string) []gocv.Mat {
	templatesInfo := *templatesMap[templateId]
	templatesInfo.usedAt = time.Now()
	return templatesInfo.templates
}

func AddTemplates(templates []gocv.Mat) string {
	templateId := uuid.New().String()
	templatesMap[templateId] = &templatesInfo{templates: templates, usedAt: time.Now()}
	return templateId
}

func init() {
	go func() {
		for {
			now := time.Now()
			time.Sleep(30 * time.Minute)
			for k, v := range templatesMap {
				if v.usedAt.Before(now) {
					delete(templatesMap, k)
				}
			}
		}
	}()
}
