package tui

import "github.com/charmbracelet/bubbles/list"

type Item struct {
	title, desc string
}

func NewItem(title, desc string) Item {
	return Item{
		title: title,
		desc:   desc,
	}
}

func (i Item) Title() string {
	return i.title
}

func (i Item) Description() string {
	return i.desc
}

func (i Item) FilterValue() string {
	return i.title
}

// Lang items
var langOptions = []list.Item{
	Item{"Go", "https://golangweekly.com/latest"},
	Item{"Python", "https://pycoders.com/latest"},
	Item{"JavaScript", "https://javascriptweekly.com/latest"},
}
