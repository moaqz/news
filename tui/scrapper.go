package tui

import (
	"log"
	"os"

	"github.com/charmbracelet/bubbles/list"
	"github.com/gocolly/colly/v2"
)

func fetchNews(lang string) []list.Item {
	var err error
	var news []list.Item

	switch lang {
	case "Go":
		news, err = getWeeklyNews("https://golangweekly.com/latest")
	case "JavaScript":
		news, err = getWeeklyNews("https://javascriptweekly.com/latest")
	case "Node.js":
		news, err = getWeeklyNews("https://nodeweekly.com//latest")
	case "Ruby":
		news, err = getWeeklyNews("https://rubyweekly.com/latest")
	case "Databases":
		news, err = getWeeklyNews("https://dbweekly.com//latest")
	case "CSS":
		news, err = getCssNews()

	default:
		log.Fatal("Invalid language")
	}

	if err != nil {
		os.Exit(1)
	}

	return news
}

// getWeeklyNews is a utility function used to extract information from the weekly news of cooperpress.
// https://cooperpress.com/publications/
func getWeeklyNews(url string) ([]list.Item, error) {
	var news []list.Item
	c := colly.NewCollector()

	c.OnHTML(".mainlink > a", func(e *colly.HTMLElement) {
		title := e.Text
		link := e.Attr("href")

		news = append(news, News{Title: title, Url: link})
	})

	if err := c.Visit(url); err != nil {
		return []list.Item{}, err
	}

	return news, nil
}

func getCssNews() ([]list.Item, error) {
	var news []list.Item
	c := colly.NewCollector()

	c.OnHTML(".archives-list article:first-child .read-more", func(e *colly.HTMLElement) {
		link := e.Attr("href")

		c.OnHTML(".article-title > a", func(e *colly.HTMLElement) {
			news = append(news, News{
				Title: e.Text,
				Url:   e.Attr("href"),
			})
		})

		if err := c.Visit(link); err != nil {
			return
		}
	})

	if err := c.Visit("https://css-weekly.com/archives"); err != nil {
		return []list.Item{}, err
	}

	return news, nil
}
