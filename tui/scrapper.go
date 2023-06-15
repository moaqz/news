package tui

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/charmbracelet/bubbles/list"
)

func getGolangNews() ([]list.Item, error) {
	r, err := http.Get("https://golangweekly.com/latest")
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

	doc.Find(".mainlink > a").Each(func(i int, s *goquery.Selection) {
		title := s.Text()
		link, _ := s.Attr("href")

		news = append(news, NewItem(title, link))
	})

	return news, nil
}

func getPythonNews() ([]list.Item, error) {
	r, err := http.Get("https://pycoders.com/latest")
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
	doc.Find(".mcnTextContent > span > a").Each(func(i int, s *goquery.Selection) {
		title := s.Text()
		link, _ := s.Attr("href")

		news = append(news, NewItem(title, link))
	})

	return news, nil
}

func getJavaScriptNews() ([]list.Item, error) {
	r, err := http.Get("https://javascriptweekly.com/latest")
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
	doc.Find(".mainlink > a").Each(func(i int, s *goquery.Selection) {
		title := s.Text()
		link, _ := s.Attr("href")

		news = append(news, NewItem(title, link))
	})

	return news, nil
}
