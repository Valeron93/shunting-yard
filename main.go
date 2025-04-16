package main

import (
	"fmt"
	"log"

	"github.com/Valeron93/shunting-yard/pkg/eval"
	"github.com/Valeron93/shunting-yard/pkg/tokenizer"
)

func main() {

	input := []rune("(2 / 10) + (1 / 5)")

	output, err := tokenizer.Tokenize(input)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tokenizer infix output: %v\n", output)
	postfix, err := tokenizer.InfixToPostfix(output)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("postfix output: %v\n", postfix)

	result, err := eval.Evaluate(postfix)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("result: %v\n", result)
}
