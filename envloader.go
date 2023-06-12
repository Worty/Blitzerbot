package main

import (
	"fmt"
	"os"
)

type Config struct {
	Twitter_target_account string
	Telegram_token         string
	Telegram_chat_id       string
}

func loadEnv() (Config, error) {
	var config Config
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
