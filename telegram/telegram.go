package telegram

import (
	"btc-news-bot/common"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SendNewsToChannel(botToken string, channelUsername string, newsItems []common.NewsItem) {
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Printf("Failed to send news to Telegram channel %s: %v", channelUsername, err)
		return // or handle the error more gracefully
	}

	for _, item := range newsItems {
		msgText := fmt.Sprintf("*%s*\n\n%s\n[Read more](%s)", item.Title, item.Description, item.Link)
		msg := tgbotapi.NewMessageToChannel(channelUsername, msgText)
		msg.ParseMode = "Markdown"
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
