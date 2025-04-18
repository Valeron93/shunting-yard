package main

import (
	"fmt"
	"strings"

	"github.com/Valeron93/shunting-yard/pkg/eval"
	"github.com/Valeron93/shunting-yard/pkg/tokenizer"
	"github.com/chzyer/readline"
)

func printWelcome() {
	fmt.Println("Shunting Yard Interactive Shell")
	fmt.Println("Type \"help\" for more information")
}

func printHelpMenu() {
	fmt.Println("Type a mathematical expression to evaluate.")
	fmt.Println("Commands:")
	fmt.Println("\thelp - show this help")
	fmt.Println("\texit - exit the program")
}

func evaluateExprAndPrint(input string) {

	tokens, err := tokenizer.Tokenize([]rune(input))

	if err != nil {
		fmt.Printf("Tokenizer failed: %v\n", err)
		return
	}

	expr, err := eval.TokensToExpression(tokens)
	if err != nil {
		fmt.Printf("Parser failed: %v\n", err)
		return
	}

	value, err := eval.Evaluate(expr)
	if err != nil {
		fmt.Printf("Evaluator failed: %v\n", err)
		return
	}

	fmt.Printf("= %v\n", value)
}

func main() {
	rl, err := readline.NewEx(&readline.Config{
		Prompt:          "> ",
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
	})
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	printWelcome()
	run := true
	for run {
		line, err := rl.Readline()
		if err != nil {
			run = false
		}

		switch line = strings.TrimSpace(line); line {
		case "exit":
			run = false

		case "help":
			printHelpMenu()

		default:
			if len(line) != 0 {
				evaluateExprAndPrint(line)
			}
		}
	}
}
