package tui

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/charmbracelet/bubbles/list"
)

var newsCache map[string][]list.Item

func getNews(url string, selector string) ([]list.Item, error) {
	// Check if the result is already in the cache
	if news, ok := newsCache[url]; ok {
		return news, nil
	}

	r, err := http.Get(url)
	if err != nil {
		return []list.Item{}, err
	}

	defer r.Body.Close()
	if r.StatusCode != 200 {
		return []list.Item{}, err
	}

	doc, err := goquery.NewDocumentFromReader(r.Body)
	if err != nil {
		return []list.Item{}, err
	}

	var news []list.Item
	doc.Find(selector).Each(func(i int, s *goquery.Selection) {
		title := s.Text()
		link, _ := s.Attr("href")

		news = append(news, News{Title: title, Url: link})
	})

	// Store the result in the cache
	newsCache[url] = news

	return news, nil
}

func getGolangNews() ([]list.Item, error) {
	return getNews("https://golangweekly.com/latest", ".mainlink > a")
}

func getJavaScriptNews() ([]list.Item, error) {
	return getNews("https://javascriptweekly.com/latest", ".mainlink > a")
}
