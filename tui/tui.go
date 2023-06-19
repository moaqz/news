package tui

import (
	"fmt"
	"os"

	"github.com/atotto/clipboard"
	"github.com/moaqz/news/ui"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type tab int

const (
	languageTab tab = iota
	newsTab
)

type Model struct {
	width        int
	height       int
	keys         KeyMap
	help         help.Model
	newsList     list.Model
	languageList list.Model
	quitting     bool
	tab          tab
}

func NewModel() Model {
	return Model{
		help:         help.New(),
		keys:         DefaultKeyMap,
		newsList:     NewNewsList(),
		languageList: NewLanguageList(),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case languageSelectionMsg:
		news := fetchNews(msg.lang)

		m.newsList.SetItems(news)
		m.newsList.Title = msg.lang + " " + "News"

	case tea.WindowSizeMsg:
		m.height = msg.Height - 3
		m.width = msg.Width

		m.newsList.SetHeight(m.height)
		m.languageList.SetHeight(m.height)

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			m.quitting = true
			return m, tea.Quit
		case key.Matches(msg, m.keys.NextTab):
			if m.newsList.SettingFilter() {
				return m, nil
			}

			m.NextTab()
		case key.Matches(msg, m.keys.PreviousTab):
			if m.newsList.SettingFilter() {
				return m, nil
			}

			m.PreviousTab()
		case key.Matches(msg, m.keys.ToggleHelp):
			m.help.ShowAll = !m.help.ShowAll

			var newHeight int
			if m.help.ShowAll {
				newHeight = m.height - 3
			} else {
				newHeight = m.height
			}

			m.newsList.SetHeight(newHeight)
			m.languageList.SetHeight(newHeight)
		case key.Matches(msg, m.keys.Search):
			m.tab = newsTab

		case key.Matches(msg, m.keys.Copy):
			if m.tab == languageTab {
				return m, nil
			}

			item := m.newsList.SelectedItem()
			if item == nil {
				return m, nil
			}

			url := item.(News).Url
			if err := clipboard.WriteAll(url); err != nil {
				return m, m.newsList.NewStatusMessage(ui.FailedMessage.Render("❌ Failed to copy to clipboard"))
			}

			return m, m.newsList.NewStatusMessage(ui.SuccessMessage.Render("🔗 Copied to clipboard!"))
		}

		switch msg.String() {
		case "enter":
			if m.tab == languageTab {
				selectedLang := m.languageList.SelectedItem().(language)
				return m, languageSelection(string(selectedLang))
			}

			if m.tab == newsTab {
				selectedItem := m.newsList.SelectedItem()

				if selectedItem == nil {
					return m, nil
				}
			}
		}
	}

	cmd := m.updateActiveTab(msg)
	return m, cmd
}

// updateActiveTab updates the currently active tab.
func (m *Model) updateActiveTab(msg tea.Msg) tea.Cmd {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch m.tab {
	case newsTab:
		m.newsList, cmd = m.newsList.Update(msg)
		cmds = append(cmds, cmd)

		m.newsList.Styles.Title = ui.FocusedTitle
		m.newsList.Styles.TitleBar = ui.FocusedTitleBar.Width(newsListWidth)

		m.languageList.Styles.Title = ui.UnFocusedTitle
		m.languageList.Styles.TitleBar = ui.UnFocusedTitleBar.Width(languageListWidth)
	case languageTab:
		m.languageList.Styles.Title = ui.FocusedTitle
		m.languageList.Styles.TitleBar = ui.FocusedTitleBar.Width(languageListWidth)

		m.newsList.Styles.Title = ui.UnFocusedTitle
		m.newsList.Styles.TitleBar = ui.UnFocusedTitleBar.Width(newsListWidth)

		m.languageList, cmd = m.languageList.Update(msg)
		cmds = append(cmds, cmd)
	}

	return tea.Batch(cmds...)
}

func (m Model) View() string {
	if m.quitting {
		return ""
	}

	return lipgloss.JoinVertical(
		lipgloss.Top,
		lipgloss.JoinHorizontal(
			lipgloss.Left,
			m.languageList.View(),
			m.newsList.View(),
		),
		ui.Margin.Render(m.help.View(m.keys)),
	)
}

// StartTea the entry point for the UI. Initializes the model.
func StartTea() {
	newsCache = make(map[string][]list.Item)
	m := NewModel()

	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

const maxTabs = 2

func (m *Model) NextTab() {
	m.tab++

	if m.tab > maxTabs {
		m.tab = 0
	}
}

func (m *Model) PreviousTab() {
	m.tab--

	if m.tab < 0 {
		m.tab = maxTabs - 1
	}
}

type languageSelectionMsg struct {
	lang string
}

func languageSelection(lang string) tea.Cmd {
	return func() tea.Msg {
		return languageSelectionMsg{lang: lang}
	}
}
