package main

import (
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func newItemDelegate(keys *delegateKeyMap) list.DefaultDelegate {
	d := list.NewDefaultDelegate()

	d.UpdateFunc = func(msg tea.Msg, m *list.Model) tea.Cmd {
		var title string
		// var moonlightCmd *exec.Cmd
		var mlItem item

		if i, ok := m.SelectedItem().(item); ok {
			mlItem = i
		} else {
			return nil
		}

		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch {
			case key.Matches(msg, keys.choose):
				type MoonlightFinishedMsg struct{ err error }
				title = mlItem.Title()
				mlConfig := mlItem.config
				moonlightExecCmd := mlConfig.GenerateCommand()
				log.Println("Generated command: ", moonlightExecCmd.Args)
				mlArgs := fmt.Sprintf("%v", moonlightExecCmd.Args)
				statusCmd := m.NewStatusMessage(statusMessageStyle("You chose " + title + " generating command: " + mlArgs))

				moonlightTeaCmd := tea.ExecProcess(moonlightExecCmd, func(err error) tea.Msg {
					return MoonlightFinishedMsg{err: err}
				})
				return tea.Sequence(statusCmd, moonlightTeaCmd)
			}
		}

		return nil
	}

	help := []key.Binding{
		keys.choose,
	}

	d.ShortHelpFunc = func() []key.Binding {
		return help
	}

	d.FullHelpFunc = func() [][]key.Binding {
		return [][]key.Binding{help}
	}

	return d
}

type delegateKeyMap struct {
	choose key.Binding
}

// Additional short help entries. This satisfies the help.KeyMap interface and
// is entirely optional.
func (d delegateKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		d.choose,
	}
}

// Additional full help entries. This satisfies the help.KeyMap interface and
// is entirely optional.
func (d delegateKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{
			d.choose,
		},
	}
}

func newDelegateKeyMap() *delegateKeyMap {
	return &delegateKeyMap{
		choose: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "choose"),
		),
	}
}
