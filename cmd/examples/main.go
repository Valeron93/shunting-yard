package main

import (
	"fmt"

	"github.com/Valeron93/shunting-yard/pkg/eval"
	"github.com/Valeron93/shunting-yard/pkg/tokenizer"
)

func main() {

	examples := []string{
		"1+1",
		"-2 + 1",
		"sin(1000)^2 + cos(1000)^2 + 10^2 + 3 mod 2",
		"2*2*2",
		"3 + 4 * 2 / (1 - 5)^2",
		"sqrt(16) + log(100)",
		"tan(45) + atan(1)",
		"abs(-42) + 7 mod 5",
		"5 + 2^3",
		"exp(1)^2 - log(7)",
		"floor(9.9) + ceil(1.1)",
		"exp(log(100))",
		"1 + sin 2",
		"pi + 1",
		"exp(2.2)",
		"e^2.2",
		"phi",
		"(1+sqrt(5))/2",
	}
	printExamples(examples)
}

func printExample(n int, example string) {

	fmt.Printf("------ EXAMPLE %v ------\n", n)
	defer fmt.Print("------------------------\n\n")
	fmt.Printf("example: \"%v\"\n", example)
	input := []rune(example)
	output, err := tokenizer.Tokenize(input)
	if err != nil {
		fmt.Printf("failed to tokenize: `%v`\n", err)
		return
	}
	fmt.Printf("tokenizer output: %v\n", output)
	expr, err := eval.TokensToExpression(output)
	if err != nil {
		fmt.Printf("failed to parse expression: `%v`\n", err)
		return
	}
	fmt.Printf("parsed expression: %v\n", expr)
	result, err := eval.Evaluate(expr)
	if err != nil {
		fmt.Printf("failed to evaluate expression: `%v`\n", err)
		return
	}
	fmt.Printf("result: %v\n", result)
}

func printExamples(examples []string) {

	for i, example := range examples {
		printExample(i+1, example)
	}
}
