package scraper

import (
	"btc-news-bot/common"
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/mmcdole/gofeed"
)

func FetchRSSNews(feedURL string) ([]common.NewsItem, error) {
	client := &http.Client{}
	request, err := http.NewRequest("GET", feedURL, nil)
	if err != nil {
		return nil, err
	}

	// Enhancing the request with more headers to mimic a web browser more closely.
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3")
	request.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	request.Header.Set("Accept-Language", "en-US,en;q=0.5")

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP error: %s", response.Status)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	fp := gofeed.NewParser()
	feed, err := fp.Parse(bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	var newsItems []common.NewsItem
	for _, item := range feed.Items {
		newsItem := common.NewsItem{
			Title:       item.Title,
			Description: item.Description,
			Link:        item.Link,
			PublishedAt: item.Published,
		}
		newsItems = append(newsItems, newsItem)
	}

	return newsItems, nil
}
