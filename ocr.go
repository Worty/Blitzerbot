package main

import (
	"fmt"
	"regexp"
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
	fmt.Printf("raw ocr: \n=======\n%s\n=======\n", text)
	lines := strings.Split(text, "\n")
	var newlines []string
	afterdate := false
	for _, line := range lines {
		if !afterdate {
			dateregex := regexp.MustCompile(`\d{2}\.\d{2}`)
			afterdate = (dateregex.MatchString(line) || strings.Contains(line, "Semistationäre Blitzanlage:"))
		}
		if afterdate {
			result := regexp.MustCompile(`ki \|`).ReplaceAllString(line, "")   // common ocr error in specific picture, remove it
			result = regexp.MustCompile(`\' a 8`).ReplaceAllString(result, "") // common ocr error in specific picture, remove it
			result = regexp.MustCompile(`ao`).ReplaceAllString(result, "")     // common ocr error in specific picture, remove it
			result = regexp.MustCompile(`N ı 8`).ReplaceAllString(result, "")  // common ocr error in specific picture, remove it
			result = regexp.MustCompile(`—`).ReplaceAllString(result, "")      // common ocr error in specific picture, remove it
			newlines = append(newlines, result)
		}
	}
	return strings.Join(newlines, "\n")
}
