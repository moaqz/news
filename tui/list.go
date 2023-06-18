package tui

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/moaqz/news/ui"
)

/* LANGUAGE LIST  */
type language string

func (l language) FilterValue() string { return "" }

type languagesDelegate struct{}

// Height is the number of lines the language list item takes up.
func (d languagesDelegate) Height() int {
	return 1
}

// Spacing is the number of lines to insert between language items.
func (d languagesDelegate) Spacing() int {
	return 0
}

// Update is what is called when the language selection is updated.
func (d languagesDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd {
	return nil
}

// Render renders a language list item.
func (d languagesDelegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
	l, ok := item.(language)
	if !ok {
		return
	}

	if index == m.Index() {
		fmt.Fprint(w, ui.SelectedTitle.Render("> "+string(l)))
		return
	}

	fmt.Fprint(w, ui.UnSelectedTitle.Render("  "+string(l)))
}

const (
	languageListWidth = 25
	newsListWidth     = 90
)

func NewLanguageList() list.Model {
	l := list.New([]list.Item{}, languagesDelegate{}, languageListWidth, 0)
	l.Title = "Languages"

	l.SetShowHelp(false)
	l.SetFilteringEnabled(false)
	l.SetShowStatusBar(false)
	l.SetStatusBarItemName("language", "languages")
	l.SetItems([]list.Item{
		language("Go"),
		language("JavaScript"),
	})

	return l
}

/* NEWS LIST */
type News struct {
	Title string
	Url   string
}

func (n News) FilterValue() string {
	return n.Title
}

type newsDelegate struct{}

func (d newsDelegate) Height() int {
	return 2
}
func (d newsDelegate) Spacing() int {
	return 1
}

func (d newsDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd {
	return nil
}

func (d newsDelegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
	if item == nil {
		return
	}

	n, ok := item.(News)
	if !ok {
		return
	}

	if index == m.Index() {
		fmt.Fprintln(w, ui.SelectedTitle.Render("> "+n.Title))
		fmt.Fprint(w, "  "+ui.SelectedSubTitle.Render(n.Url))
		return
	}

	fmt.Fprintln(w, "  "+ui.UnSelectedTitle.Render(n.Title))
	fmt.Fprintf(w, "  "+ui.UnSelectedSubTitle.Render(n.Url))
}

func NewNewsList() list.Model {
	l := list.New([]list.Item{}, newsDelegate{}, newsListWidth, 0)
	l.Title = "News"
	l.SetShowHelp(false)
	l.SetStatusBarItemName("news", "news")
	l.Styles.NoItems = ui.NoItems
	l.SetShowStatusBar(false)

	return l
}
