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
	Green   = lipgloss.Color("#a6e3a1")
	Red     = lipgloss.Color("#f38ba8")
)

var (
	Margin             = lipgloss.NewStyle().Margin(1, 0, 0, 1)
	NoItems            = lipgloss.NewStyle().Margin(0, 2).Foreground(Surface)
	SelectedTitle      = lipgloss.NewStyle().Foreground(Muave).Bold(true)
	SelectedSubTitle   = lipgloss.NewStyle().Foreground(SubText)
	UnSelectedTitle    = lipgloss.NewStyle().Foreground(Surface).Bold(true)
	UnSelectedSubTitle = lipgloss.NewStyle().Foreground(Base)
	FocusedTitle       = lipgloss.NewStyle().Foreground(Peach).Padding(0, 2)
	UnFocusedTitle     = lipgloss.NewStyle().Foreground(Surface).Padding(0, 2)
	FocusedTitleBar    = lipgloss.NewStyle().Margin(0, 1, 1, 1).Border(lipgloss.DoubleBorder()).BorderForeground(Peach)
	UnFocusedTitleBar  = lipgloss.NewStyle().Margin(0, 1, 1, 1).Border(lipgloss.DoubleBorder()).BorderForeground(Surface)
	Cursor             = lipgloss.NewStyle().Foreground(Peach)
	Prompt             = lipgloss.NewStyle().Padding(0, 2).Foreground(Peach)
	SuccessMessage     = lipgloss.NewStyle().Foreground(Green)
	FailedMessage      = lipgloss.NewStyle().Foreground(Red)
)
