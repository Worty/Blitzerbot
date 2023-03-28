package main

import (
	"strings"

	"github.com/otiai10/gosseract/v2"
)

func doOCR(url string) (string, error) {
	path, err := downloadImage(url)
	if err != nil {
		return "", err
	}
	client := gosseract.NewClient()
	defer client.Close()
	client.SetLanguage("deu")
	client.SetImage(path)
	text, _ := client.Text()
	return text, nil
}

func cleanupOCRText(text string) string {
	lines := strings.Split(text, "\n")
	var newlines []string
	for _, line := range lines {
		if strings.Contains(line, ":") {
			newlines = append(newlines, line)
		}
	}
	return strings.Join(newlines, "\n")
}
