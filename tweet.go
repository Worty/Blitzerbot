package main

import (
	"strings"
	"time"

	"github.com/cavaliergopher/grab/v3"
)

var keywords = []string{"blitzen", "Geschwindigkeitskontrollen", "Semi-Station", "geblitzt"}

func checkIfKeywordInText(text string) bool {
	for _, keyword := range keywords {
		if strings.Contains(text, keyword) {
			return true
		}
	}
	return false
}

func checkIfInTargetTimeRange(date time.Time) bool {
	yesterday := time.Now().AddDate(0, 0, -1)
	targetbegin := time.Date(yesterday.Year(), yesterday.Month(), yesterday.Day(), 16, 55, 0, 0, time.Local)
	targetend := time.Date(yesterday.Year(), yesterday.Month(), yesterday.Day(), 17, 15, 0, 0, time.Local)
	return date.After(targetbegin) && date.Before(targetend)
}

func downloadImage(url string) (string, error) {
	resp, err := grab.Get("/tmp", url)
	if err != nil {
		return "", err
	}
	return resp.Filename, nil
}
