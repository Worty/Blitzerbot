package main

import (
	"fmt"
	"os"
)

type Config struct {
	Twitter_consumer_key    string
	Twitter_consumer_secret string
	Twitter_access_token    string
	Twitter_access_secret   string
	Twitter_target_account  string
	Telegram_token          string
	Telegram_chat_id        string
}

func loadEnv() (Config, error) {
	var config Config
	config.Twitter_consumer_key = os.Getenv("TWITTER_CONSUMER_KEY")
	if config.Twitter_consumer_key == "" {
		return config, fmt.Errorf("TWITTER_CONSUMER_KEY not set")
	}
	config.Twitter_consumer_secret = os.Getenv("TWITTER_CONSUMER_SECRET")
	if config.Twitter_consumer_secret == "" {
		return config, fmt.Errorf("TWITTER_CONSUMER_SECRET not set")
	}
	config.Twitter_access_token = os.Getenv("TWITTER_ACCESS_TOKEN")
	if config.Twitter_access_token == "" {
		return config, fmt.Errorf("TWITTER_ACCESS_TOKEN not set")
	}
	config.Twitter_access_secret = os.Getenv("TWITTER_ACCESS_SECRET")
	if config.Twitter_access_secret == "" {
		return config, fmt.Errorf("TWITTER_ACCESS_SECRET not set")
	}
	config.Twitter_target_account = os.Getenv("TWITTER_TARGET_ACCOUNT")
	if config.Twitter_target_account == "" {
		return config, fmt.Errorf("TWITTER_TARGET_ACCOUNT not set")
	}
	config.Telegram_token = os.Getenv("TELEGRAM_TOKEN")
	if config.Telegram_token == "" {
		return config, fmt.Errorf("TELEGRAM_TOKEN not set")
	}
	config.Telegram_chat_id = os.Getenv("TELEGRAM_CHAT_ID")
	if config.Telegram_chat_id == "" {
		return config, fmt.Errorf("TELEGRAM_CHAT_ID not set")
	}
	return config, nil
}
