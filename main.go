package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dghubble/oauth1"
	twitter "github.com/g8rswimmer/go-twitter/v2"
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
	config := oauth1.NewConfig(envConfig.Twitter_consumer_key, envConfig.Twitter_consumer_secret)
	token := oauth1.NewToken(envConfig.Twitter_access_token, envConfig.Twitter_access_secret)
	httpClient := config.Client(oauth1.NoContext, token)

	client := &twitter.Client{
		Authorizer: &authorize{},
		Client:     httpClient,
		Host:       "https://api.twitter.com",
	}

	opts := twitter.UserTweetTimelineOpts{
		TweetFields: []twitter.TweetField{twitter.TweetFieldCreatedAt, twitter.TweetFieldAuthorID, twitter.TweetFieldConversationID, twitter.TweetFieldAttachments},
		Expansions:  []twitter.Expansion{twitter.ExpansionAttachmentsMediaKeys},
		UserFields:  []twitter.UserField{},
		MediaFields: []twitter.MediaField{twitter.MediaFieldURL, twitter.MediaFieldPreviewImageURL, twitter.MediaFieldType},
		MaxResults:  5,
	}

	log.Println("Getting timeline...")
	timeline, err := client.UserTweetTimeline(context.Background(), envConfig.Twitter_target_account, opts)
	if err != nil {
		log.Panicf("user tweet timeline error: %v", err)
	}

	for _, tweet := range timeline.Raw.Tweets {
		if strings.Contains(tweet.Text, "blitzen") || strings.Contains(tweet.Text, "Geschwindigkeitskontrollen") {
			tweetdate, err := time.Parse(time.RFC3339, tweet.CreatedAt)
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

			if tweet.Attachments == nil || len(tweet.Attachments.MediaKeys) == 0 {
				log.Println("found Blitzer in time but no image")
				continue
			}

			mediakey := tweet.Attachments.MediaKeys[0]
			url, err := findUrlByMediaKey(mediakey, timeline.Raw.Includes.Media)
			if err != nil {
				log.Printf("found Blitzer but no image: %v\n", err)
				continue
			}

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
