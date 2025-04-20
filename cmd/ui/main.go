package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/Valeron93/shunting-yard/pkg/eval"
	"github.com/Valeron93/shunting-yard/pkg/tokenizer"
	"github.com/chzyer/readline"
	"github.com/fatih/color"
)

var verbose bool = false

func printWelcome() {
	color.Yellow("Shunting Yard Interactive Shell")
	color.HiWhite("Type \"help\" for more information")
	if verbose {
		color.Green("Running in verbose mode")
	}
}

func printHelpMenu() {
	c := color.New(color.Underline)

	c.Println("Type a mathematical expression to evaluate.")
	fmt.Println("Commands:")
	fmt.Printf("\t%s - show this help\n", color.GreenString("help"))
	fmt.Printf("\t%s - exit the program\n", color.GreenString("exit"))
	fmt.Println("Flags:")
	fmt.Printf("\t%s - show additional tokenizer and parser output\n", color.GreenString("-verbose"))
}

func evaluateExprAndPrint(input string) {

	tokens := tokenizer.Tokenize(input)

	if verbose {
		color.Green("Tokens: %v", tokens)
	}

	expr, err := eval.TokensToExpression(tokens)
	if err != nil {
		color.Red("Parser failed: %v\n", err)
		return
	}

	if verbose {
		color.Green("Expression: %v", expr)
	}

	value, err := eval.Evaluate(expr)
	if err != nil {
		color.Red("Evaluator failed: %v\n", err)
		return
	}

	fmt.Printf("%s %v\n", color.YellowString("="), value)
}

func main() {
	flag.BoolVar(&verbose, "verbose", false, "show tokenizer and parser output")
	flag.Parse()

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
