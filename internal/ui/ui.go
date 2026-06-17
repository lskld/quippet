package ui

import (
	"strings"

	"github.com/lskld/quippet/internal/model"

	tea "github.com/charmbracelet/bubbletea"
)

type listModel struct {
	snippets []model.Snippet
}

func (m listModel) Init() tea.Cmd {
	return nil
}
func (m listModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" {
			return m, tea.Quit
		}
	}
	return m, nil
}
func (m listModel) View() string {
	var sb strings.Builder

	sb.WriteString("Quippet\n\n")

	for _, snippet := range m.snippets {
		sb.WriteString(snippet.Title)
		sb.WriteByte('\n')
	}

	sb.WriteString("\npress q to quit")

	return sb.String()
}

func Run(snippets []model.Snippet) error {
	m := listModel{snippets: snippets}
	p := tea.NewProgram(m)
	_, err := p.Run()
	return err
}