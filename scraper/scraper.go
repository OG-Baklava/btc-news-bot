package scraper

import (
	"btc-news-bot/common"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mmcdole/gofeed"
)

func FetchRSSNews(feedURL string) ([]common.NewsItem, error) {
	client := &http.Client{}
	request, err := http.NewRequest("GET", feedURL, nil)
	if err != nil {
		return nil, err
	}

	// Setting a user-agent header
	request.Header.Set("User-Agent", "Mozilla/5.0")

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP error: %s", response.Status)
	}

	body, err := ioutil.ReadAll(response.Body)
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
