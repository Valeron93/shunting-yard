package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Valeron93/shunting-yard/pkg/eval"
	"github.com/Valeron93/shunting-yard/pkg/tokenizer"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

// TODO: add history, maybe even store it in a file

// TODO: we should find a way to write program output to stdout another
// way, storing every string is definitely not optimal
type model struct {
	input  textinput.Model
	output []string
}

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = ""
	ti.Focus()
	ti.CharLimit = 256
	ti.Width = 50
	ti.Prompt = " > "

	return model{
		input:  ti,
		output: []string{"Shunting Yard Evaluator"},
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func parseExpressionCommand(input string) (result float64, err error) {
	tokens, err := tokenizer.Tokenize([]rune(input))
	if err != nil {
		return
	}

	expr, err := eval.TokensToExpression(tokens)
	if err != nil {
		return
	}
	result, err = eval.Evaluate(expr)
	return
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			command := strings.TrimSpace(m.input.Value())
			m.output = append(m.output, fmt.Sprintf("%s%s", m.input.Prompt, command))

			switch command {
			case "exit":
				return m, tea.Quit

			case "help":
				m.output = append(m.output, " : type a mathematical expression to evaluate")
				m.output = append(m.output, " : `exit`, ^D or ^C to exit")

			default:
				if result, err := parseExpressionCommand(command); err != nil {
					m.output = append(m.output, fmt.Sprintf("parsing error: %v", err))
				} else {
					m.output = append(m.output, fmt.Sprintf(" = %v", result))
				}
			}
			m.input.SetValue("")

		case tea.KeyCtrlC, tea.KeyCtrlD:
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		m.input.Width = msg.Width
	}

	m.input, cmd = m.input.Update(msg)
	return m, cmd
}

func (m model) View() string {
	s := strings.Join(m.output, "\n")
	s += "\n" + m.input.View()
	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
