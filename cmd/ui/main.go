package main

import (
	"fmt"
	"strings"

	"github.com/Valeron93/shunting-yard/pkg/eval"
	"github.com/Valeron93/shunting-yard/pkg/tokenizer"
	"github.com/chzyer/readline"
	"github.com/fatih/color"
)

func printWelcome() {
	color.Yellow("Shunting Yard Interactive Shell")
	color.HiWhite("Type \"help\" for more information")
}

func printHelpMenu() {
	c := color.New(color.Underline)

	c.Println("Type a mathematical expression to evaluate.")
	fmt.Println("Commands:")
	fmt.Printf("\t%s - show this help\n", color.GreenString("help"))
	fmt.Printf("\t%s - exit the program\n", color.GreenString("exit"))
}

func evaluateExprAndPrint(input string) {

	tokens, err := tokenizer.Tokenize([]rune(input))

	if err != nil {
		color.Red("Tokenizer failed: %v\n", err)
		return
	}

	expr, err := eval.TokensToExpression(tokens)
	if err != nil {
		color.Red("Parser failed: %v\n", err)
		return
	}

	value, err := eval.Evaluate(expr)
	if err != nil {
		color.Red("Evaluator failed: %v\n", err)
		return
	}

	fmt.Printf("%s %v\n", color.YellowString("="), value)
}

func main() {
	rl, err := readline.NewEx(&readline.Config{
		Prompt:          color.YellowString("> "),
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
	})
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	printWelcome()
	for {
		line, err := rl.Readline()
		if err != nil {
			return
		}

		switch line = strings.TrimSpace(line); line {
		case "exit":
			return

		case "help":
			printHelpMenu()

		default:
			if len(line) != 0 {
				evaluateExprAndPrint(line)
			}
		}
	}
}
