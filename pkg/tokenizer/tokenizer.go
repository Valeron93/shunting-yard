package tokenizer

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"

	"github.com/Valeron93/shunting-yard/pkg/operator"
	"github.com/Valeron93/shunting-yard/pkg/stack"
)

type TokenType int16

const (
	Op     TokenType = iota
	Number TokenType = iota
)

type Token struct {
	Type TokenType
	Data any
}

func parseNumber(input []rune) (int, float64, error) {
	var i int
	for _, c := range input {
		if unicode.IsNumber(c) || c == '.' {
			i++
		} else {
			break
		}
	}

	num, err := strconv.ParseFloat(string(input[0:i]), 64)
	return i, num, err
}

func (t Token) String() string {

	switch t.Type {

	case Op:
		op := string(t.Data.(operator.Operator))
		return fmt.Sprintf("%v", op)

	case Number:
		return fmt.Sprintf("%v", t.Data.(float64))
	default:
		return "unknown token"
	}

}

func Tokenize(input []rune) ([]Token, error) {
	tokens := make([]Token, 0, 10)

	i := 0
	for i < len(input) {
		c := input[i]

		if unicode.IsSpace(c) {
			i++
			continue
		}
		if unicode.IsNumber(c) || c == '.' {
			count, num, err := parseNumber(input[i:])
			if err != nil {
				return nil, err
			}

			tokens = append(tokens, Token{
				Type: Number,
				Data: num,
			})

			i += count
			continue
		}
		op, err := operator.FromRune(c)
		if err != nil {
			return nil, err
		}
		tokens = append(tokens, Token{
			Type: Op,
			Data: op,
		})
		i++
	}

	return tokens, nil
}

func InfixToPostfix(infix []Token) ([]Token, error) {
	postfix := make([]Token, 0, 10)
	var stack stack.Stack[operator.Operator]

	for _, token := range infix {

		if token.Type == Number {
			postfix = append(postfix, token)
			continue
		}

		op := token.Data.(operator.Operator)

		if op == operator.ParenOpen {
			stack.Push(op)
			continue
		}

		if op == operator.ParenClose {
			for stack.Count() > 0 && stack.MustPeek() != operator.ParenOpen {
				postfix = append(postfix, NewTokenFromOperator(stack.MustPop()))
			}
			_, ok := stack.Pop()
			if !ok {
				return nil, errors.New("stack was empty")
			}
			continue
		}

		for stack.Count() > 0 && op.Precedence() <= stack.MustPeek().Precedence() {
			postfix = append(postfix, NewTokenFromOperator(stack.MustPop()))
		}
		stack.Push(op)
	}

	for stack.Count() > 0 {
		postfix = append(postfix, NewTokenFromOperator(stack.MustPop()))
	}

	return postfix, nil
}

func NewTokenFromOperator(op operator.Operator) Token {
	return Token{
		Type: Op,
		Data: op,
	}
}
