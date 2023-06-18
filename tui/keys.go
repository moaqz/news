package tui

import "github.com/charmbracelet/bubbles/key"

// KeyMap is the mappings of actions to key bindings.
type KeyMap struct {
	Quit        key.Binding
	Search      key.Binding
	ToggleHelp  key.Binding
	MoveUp      key.Binding
	MoveDown    key.Binding
	NextTab     key.Binding
	PreviousTab key.Binding
}

// DefaultKeyMap is the default key map for the application.
var DefaultKeyMap = KeyMap{
	Quit:        key.NewBinding(key.WithKeys("q", "ctrl+c"), key.WithHelp("q", "exit")),
	Search:      key.NewBinding(key.WithKeys("/"), key.WithHelp("/", "search")),
	ToggleHelp:  key.NewBinding(key.WithKeys("?"), key.WithHelp("?", "help")),
	MoveUp:      key.NewBinding(key.WithKeys("j"), key.WithHelp("j", "move down")),
	MoveDown:    key.NewBinding(key.WithKeys("k"), key.WithHelp("k", "move up")),
	NextTab:     key.NewBinding(key.WithKeys("tab", "right"), key.WithHelp("tab", "navigate")),
	PreviousTab: key.NewBinding(key.WithKeys("shift+tab", "left"), key.WithHelp("shift+tab", "navigate")),
}

// ShortHelp returns a quick help menu.
func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Search,
		k.NextTab,
		k.Quit,
		k.ToggleHelp,
	}
}

// FullHelp returns all help options in a more detailed view.
func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.MoveUp, k.MoveDown, k.NextTab, k.PreviousTab},
		{k.Search, k.ToggleHelp, k.Quit},
	}
}
