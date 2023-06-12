package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"time"

	twitterscraper "github.com/n0madic/twitter-scraper"
)

type authorize struct{}

func (a authorize) Add(req *http.Request) {}

var envConfig Config

func main() {
	loc, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		log.Panicf("error loading location: %v", err)
	}
	envConfig, err = loadEnv()
	if err != nil {
		log.Panicf("error loading env: %v", err)
	}

	prodmode := flag.Bool("prodmode", false, "If true, send via telegram")
	flag.Parse()
	// init twitter client
	scraper := twitterscraper.New()

	log.Println("Getting timeline...")
	ch := scraper.GetTweets(context.Background(), envConfig.Twitter_target_account, 10)

	for tweet := range ch {
		if checkIfKeywordInText(tweet.Text) {
			tweetdate := tweet.TimeParsed
			if err != nil {
				log.Printf("error parsing date: %v\n", err)
				continue
			}
			tweetdate = tweetdate.In(loc)
			log.Printf("[%s]: %s\n", tweetdate.Format(time.RFC3339), tweet.Text)
			if !checkIfInTargetTimeRange(tweetdate) {
				log.Println("Not in target time range")
				continue
			}

			if tweet.Photos == nil || len(tweet.Photos) == 0 {
				log.Println("found in time but no image")
				continue
			}

			url := tweet.Photos[0].URL
			log.Printf("URL: %s\n", url)
			text, err := doOCR(url)
			if err != nil {
				log.Printf("error doing OCR: %v\n", err)
				continue
			}

			cleanuptext := cleanupOCRText(text)
			if *prodmode {
				if err := sendViaTelegram(cleanuptext); err != nil {
					log.Printf("error sending via telegram:\n %v\n", err)
				}
			} else {
				log.Printf("OCR Text: %s\n", cleanuptext)
			}
			return
		}
	}
}
