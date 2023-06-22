package tui

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

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
	technologyTab tab = iota
	newsTab
)

type Model struct {
	width          int
	height         int
	keys           KeyMap
	help           help.Model
	newsList       list.Model
	technologyList list.Model
	quitting       bool
	tab            tab
}

func NewModel() Model {
	return Model{
		help:           help.New(),
		keys:           DefaultKeyMap,
		newsList:       NewNewsList(),
		technologyList: NewLanguageList(),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case technologySelectionMsg:
		news := fetchNews(msg.lang)

		m.newsList.SetItems(news)
		m.newsList.Title = msg.lang + " " + "News"

	case tea.WindowSizeMsg:
		m.height = msg.Height - 3
		m.width = msg.Width

		m.newsList.SetHeight(m.height)
		m.technologyList.SetHeight(m.height)

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

			m.newsList.SetHeight(m.height)
			m.technologyList.SetHeight(m.height)
		case key.Matches(msg, m.keys.Search):
			m.tab = newsTab

		case key.Matches(msg, m.keys.Copy):
			if m.tab == technologyTab {
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
			if m.tab == technologyTab {
				selectedLang := string(m.technologyList.SelectedItem().(technology))

				statusCmd := m.technologyList.NewStatusMessage(ui.SuccessMessage.Render("Fetching news..."))
				langCmd := technologySelection(selectedLang)

				return m, tea.Batch(statusCmd, langCmd)
			}

			if m.tab == newsTab {
				item := m.newsList.SelectedItem()
				if item == nil {
					return m, nil
				}

				url := item.(News).Url
				if err := openBrowser(url); err != nil {
					m.newsList.NewStatusMessage(ui.FailedMessage.Render("❌ Failed to open url"))
				}
			}
		}
	}

	m.updateKeyMap()
	return m, m.updateActiveTab(msg)
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

		m.technologyList.Styles.Title = ui.UnFocusedTitle
		m.technologyList.Styles.TitleBar = ui.UnFocusedTitleBar.Width(technologyListWidth)
	case technologyTab:
		m.technologyList.Styles.Title = ui.FocusedTitle
		m.technologyList.Styles.TitleBar = ui.FocusedTitleBar.Width(technologyListWidth)

		m.newsList.Styles.Title = ui.UnFocusedTitle
		m.newsList.Styles.TitleBar = ui.UnFocusedTitleBar.Width(newsListWidth)

		m.technologyList, cmd = m.technologyList.Update(msg)
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
			m.technologyList.View(),
			m.newsList.View(),
		),
		ui.Margin.Render(m.help.View(m.keys)),
	)
}

// StartTea the entry point for the UI. Initializes the model.
func StartTea() {
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

type technologySelectionMsg struct {
	lang string
}

func technologySelection(lang string) tea.Cmd {
	return func() tea.Msg {
		return technologySelectionMsg{lang: lang}
	}
}

// updateKeyMap disables or enables the keys based on the current tab.
func (m *Model) updateKeyMap() {
	m.keys.Copy.SetEnabled(m.tab == newsTab && m.newsList.SelectedItem() != nil)
}

func openBrowser(targetURL string) error {
	var err error

	switch runtime.GOOS {
	case "linux":
		if os.Getenv("WSL_DISTRO_NAME") != "" {
			err = exec.Command("wslview", targetURL).Start()
		} else {
			err = exec.Command("xdg-open", targetURL).Start()
		}
	case "windows":
		err = exec.Command("rundll32.exe", "url.dll,FileProtocolHandler", targetURL).Start()
	case "darwin":
		err = exec.Command("open", targetURL).Start()
	default:
		err = fmt.Errorf("unsupported platform %v", runtime.GOOS)
	}

	return err
}
