package tui

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/moaqz/news/tui/entryui"
	"github.com/moaqz/news/tui/newsui"
	"github.com/moaqz/news/tui/shared"
)

type sessionState int

const (
	entryView sessionState = iota
	newsView
)

type MainModel struct {
	state sessionState
	entry tea.Model
	news  tea.Model
}

// New initialize the main model for your program
func NewMainModel() MainModel {
	return MainModel{
		state: sessionState(0),
		entry: entryui.NewEntryModel(),
		news:  newsui.NewNewsModel(),
	}
}

// StartTea the entry point for the UI. Initializes the model.
func StartTea() {
	m := NewMainModel()

	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

// Init run any intial IO on program start
func (m MainModel) Init() tea.Cmd {
	return nil
}

// Update handle IO and commands
func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		shared.WindowSize = msg
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	switch m.state {
	case entryView:
		newEntry, newCmd := m.entry.Update(msg)
		entryModel, ok := newEntry.(entryui.EntryModel)

		if !ok {
			panic("could not perform assertion on entryui model")
		}

		m.entry = entryModel
		cmd = newCmd
	case newsView:
		// TODO: implement news View
		panic("Not implemented")
	}

	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

// View return the text UI to be output to the terminal
func (m MainModel) View() string {
	switch m.state {
	case newsView:
		return m.news.View()
	default:
		return m.entry.View()
	}
}
