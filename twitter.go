package main

import (
	"fmt"
	"time"

	"github.com/cavaliergopher/grab/v3"
	twitter "github.com/g8rswimmer/go-twitter/v2"
)

func findUrlByMediaKey(mediaKey string, media []*twitter.MediaObj) (string, error) {
	for _, m := range media {
		if m.Key == mediaKey && m.Type == "photo" {
			return m.URL, nil
		}
	}
	return "", fmt.Errorf("could not find media key %s", mediaKey)
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
