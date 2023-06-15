package tui

import (
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type status int

const (
	lang status = iota
	news
)

var (
	divisor      = 2
	columnStyle  = lipgloss.NewStyle().Padding(1, 2).Border(lipgloss.HiddenBorder())
	focusedStyle = lipgloss.NewStyle().Padding(1, 2).Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("62"))
)

type Model struct {
	lists    []list.Model
	quitting bool
	focused  status
}

func NewModel() Model {
	// Init lang list
	langList := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	langList.Title = "Languages"
	langList.SetItems(langOptions)
	langList.SetShowHelp(false)

	// Init news list
	newsList := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	newsList.Title = "News"
	newsList.SetShowHelp(false)

	return Model{
		lists: []list.Model{langList, newsList},
	}
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

func (m Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Next() {
	if m.focused == lang {
		m.focused = news
	} else {
		m.focused = lang
	}
}

func (m *Model) Prev() {
	if m.focused == lang {
		m.focused = news
	} else {
		m.focused--
	}
}

func (m Model) FetchNews(lang string) {
	var (
		err      error
		newsList []list.Item
	)

	switch lang {
	case "Go":
		newsList, err = getGolangNews()

	case "Python":
		newsList, err = getPythonNews()

	case "JavaScript":
		newsList, err = getJavaScriptNews()
	}

	if err != nil {
		log.Fatal(err)
	}

	m.lists[news].SetItems(newsList)
	m.lists[news].Title = fmt.Sprintf("%s News", lang)
}

func OpenUrl(url string) {
	log.Fatal("not implemented")
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		columnWidth := msg.Width / divisor
		columnHeight := msg.Height - divisor

		columnStyle.Width(columnWidth)
		focusedStyle.Width(columnWidth)
		columnStyle.Height(columnHeight)
		focusedStyle.Height(columnHeight)

		for i, list := range m.lists {
			list.SetSize(msg.Width/divisor, msg.Height/divisor)
			m.lists[i], _ = list.Update(msg)
		}

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit
		case "right", "l":
			m.Next()
		case "left", "h":
			m.Prev()
		case "enter":
			if m.focused == lang {
				selectedItem := m.lists[lang].SelectedItem()
				selectedLang := selectedItem.(Item)
				m.FetchNews(selectedLang.title)

				m.Next()

				return m, nil
			}

			if m.focused == news {
				selectedItem := m.lists[news].SelectedItem()

				if selectedItem == nil {
					return m, nil
				}

				selectedNews := selectedItem.(Item)
				url := selectedNews.desc

				OpenUrl(url)
			}
		}
	}

	var cmd tea.Cmd
	m.lists[m.focused], cmd = m.lists[m.focused].Update(msg)
	return m, cmd
}

func (m Model) View() string {
	if m.quitting {
		return ""
	}

	langView := m.lists[lang].View()
	newsView := m.lists[news].View()

	switch m.focused {
	case news:
		return lipgloss.JoinHorizontal(
			lipgloss.Left,
			columnStyle.Render(langView),
			focusedStyle.Render(newsView),
		)
	default:
		return lipgloss.JoinHorizontal(
			lipgloss.Left,
			focusedStyle.Render(langView),
			columnStyle.Render(newsView),
		)
	}
}
