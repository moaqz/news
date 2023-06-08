package shared

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func LogMsg(msg tea.Msg) {
	file, err := os.OpenFile("keys.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic("RIP.")
	}
	defer file.Close()

	defaultMsg := fmt.Sprintf("Default message received: %v\n", msg)
	_, err = file.WriteString(defaultMsg)
	if err != nil {
		panic("RIP.")
	}
}
