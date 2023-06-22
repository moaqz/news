package tui

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/moaqz/news/ui"
)

type technology string

func (l technology) FilterValue() string { return "" }

type technologiesDelegate struct{}

func (d technologiesDelegate) Height() int {
	return 1
}

func (d technologiesDelegate) Spacing() int {
	return 0
}

func (d technologiesDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd {
	return nil
}

func (d technologiesDelegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
	l, ok := item.(technology)
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
	technologyListWidth = 40
	newsListWidth       = 90
)

var langListOptions = []list.Item{
	technology("Go"),
	technology("JavaScript"),
	technology("Node.js"),
	technology("Ruby"),
	technology("Databases"),
	technology("CSS"),
}

func NewLanguageList() list.Model {
	l := list.New(langListOptions, technologiesDelegate{}, technologyListWidth, 0)
	l.Title = "Technologies"

	l.SetShowHelp(false)
	l.SetFilteringEnabled(false)
	l.SetShowStatusBar(false)
	l.SetStatusBarItemName("technology", "technologies")

	return l
}

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
	l.FilterInput.PromptStyle = ui.Prompt
	l.FilterInput.Cursor.Style = ui.Cursor

	return l
}
