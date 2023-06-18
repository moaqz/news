package ui

import (
	"github.com/charmbracelet/lipgloss"
)

const (
	White   = lipgloss.Color("#FFFFFF")
	Peach   = lipgloss.Color("#F5A97F")
	Crust   = lipgloss.Color("#181926")
	Muave   = lipgloss.Color("#c6a0f6")
	SubText = lipgloss.Color("#a5adcb")
	Surface = lipgloss.Color("#494d64")
	Base    = lipgloss.Color("#363a4f")
)

var (
	Margin             = lipgloss.NewStyle().Margin(1, 0, 0, 1)
	NoItems            = lipgloss.NewStyle().Margin(0, 2).Foreground(Surface)
	SelectedTitle      = lipgloss.NewStyle().Foreground(Muave).Bold(true)
	SelectedSubTitle   = lipgloss.NewStyle().Foreground(SubText)
	UnSelectedTitle    = lipgloss.NewStyle().Foreground(Surface).Bold(true)
	UnSelectedSubTitle = lipgloss.NewStyle().Foreground(Base)
	FocusedTitle       = lipgloss.NewStyle().Foreground(Crust).Padding(0, 2)
	UnFocusedTitle     = lipgloss.NewStyle().Foreground(White).Padding(0, 2)
	FocusedTitleBar    = lipgloss.NewStyle().Background(Peach).Margin(0, 1, 1, 1)
	UnFocusedTitleBar  = lipgloss.NewStyle().Background(Surface).Margin(0, 1, 1, 1)
	Cursor             = lipgloss.NewStyle().Foreground(Crust)
	Prompt             = lipgloss.NewStyle().Foreground(Crust).Padding(0, 2)
)
