package tui

import "github.com/charmbracelet/bubbles/key"

// KeyMap is the mappings of actions to key bindings.
type KeyMap struct {
	Quit        key.Binding
	Search      key.Binding
	ToggleHelp  key.Binding
	MoveUp      key.Binding
	MoveDown    key.Binding
	GoToEnd     key.Binding
	GoToStart   key.Binding
	NextTab     key.Binding
	PreviousTab key.Binding
	Copy        key.Binding
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
	Copy:        key.NewBinding(key.WithKeys("c"), key.WithHelp("c", "copy url")),
	GoToEnd:     key.NewBinding(key.WithKeys("end", "G"), key.WithHelp("G/end", "go to end")),
	GoToStart:   key.NewBinding(key.WithKeys("home", "g"), key.WithHelp("g/home", "go to start")),
}

// ShortHelp returns a quick help menu.
func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Search,
		k.NextTab,
		k.Copy,
		k.Quit,
		k.ToggleHelp,
	}
}

// FullHelp returns all help options in a more detailed view.
func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.MoveUp, k.MoveDown},
		{k.Search, k.Copy},
		{k.GoToEnd, k.GoToStart},
		{k.NextTab, k.PreviousTab},
		{k.Quit, k.ToggleHelp},
	}
}
