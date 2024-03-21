package main

import (
	"btc-news-bot/scraper"
	"btc-news-bot/telegram"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile) // Log the timestamp and file location

	err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	channelUsername := os.Getenv("TELEGRAM_CHANNEL_USERNAME")
	rssFeedUrl := "https://cointelegraph.com/rss/tag/bitcoin"

	newsItems, err := scraper.FetchRSSNews(rssFeedUrl)
	if err != nil {
		log.Printf("Failed to fetch RSS news from %s: %v", rssFeedUrl, err)
		return // or handle the error more gracefully
	}

	telegram.SendNewsToChannel(botToken, channelUsername, newsItems)
}