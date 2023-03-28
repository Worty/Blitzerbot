package main

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func sendViaTelegram(msg string) error {
	tg, err := tgbotapi.NewBotAPI(envConfig.Telegram_token)
	channel := envConfig.Telegram_chat_id
	if err != nil {
		return err
	}
	newmsg := tgbotapi.NewMessageToChannel(channel, msg)
	newmsg.DisableNotification = true
	newmsg.DisableWebPagePreview = true
	_, err = tg.Send(newmsg)
	return err
}
