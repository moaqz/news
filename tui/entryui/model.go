package entryui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/moaqz/news/tui/shared"
)

var entryList = []list.Item{
	shared.NewListItem("Go", "https://golangweekly.com/latest"),
	shared.NewListItem("Python", "https://pycoders.com/latest"),
	shared.NewListItem("JavaScript", "https://javascriptweekly.com/latest"),
}

type EntryModel struct {
	list     list.Model
	quitting bool
}

func NewEntryModel() EntryModel {
	l := list.New(entryList, list.NewDefaultDelegate(), 0, 0)
	l.Title = "News - Ultimate Cli to watch tech novelty"

	return EntryModel{list: l}
}

func (m EntryModel) Init() tea.Cmd {
	return nil
}

func (m EntryModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetSize(msg.Width, msg.Height)
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			// TODO: fetch the news of the selected language
			// and change the view.
			return m, tea.Quit
		}
	}

	shared.LogMsg(msg)
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m EntryModel) View() string {
	if m.quitting {
		return "\nSee you later!\n\n"
	}

	return m.list.View()
}
