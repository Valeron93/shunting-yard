package eval

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/Valeron93/shunting-yard/pkg/stack"
	"github.com/Valeron93/shunting-yard/pkg/tokenizer"
)

type expressionNode struct {
	data any
}

func (e expressionNode) String() string {
	return fmt.Sprintf("%v", e.data)
}

type Expression []expressionNode

func TokensToExpression(tokens []tokenizer.Token) (Expression, error) {
	postfix := make(Expression, 0, 10)
	var stack stack.Stack[operator]

	pushFunc := func(f operator) {
		for stack.Count() > 0 && f.Precedence() <= stack.MustPeek().Precedence() {
			postfix = append(postfix, expressionNode{
				data: stack.MustPop(),
			})
		}
		stack.Push(f)
	}

	for _, token := range tokens {

		switch token.Type {

		case tokenizer.Number:
			num, err := strconv.ParseFloat(token.Str, 64)
			if err != nil {
				return postfix, err
			}
			postfix = append(postfix, expressionNode{
				data: num,
			})

		case tokenizer.Identifier:
			if f, ok := defaultFunctionMap[token.Str]; ok {
				pushFunc(f)
				continue
			}

			if num, ok := defaultConstantMap[token.Str]; ok {
				postfix = append(postfix, expressionNode{
					data: num,
				})
				continue
			}

			return postfix, fmt.Errorf("unknown identifier: %v", token.Str)

		case tokenizer.Operator:
			op, ok := defaultOperatorMap[token.Str]
			if !ok {
				return postfix, fmt.Errorf("unknown operator %v", token.Str)
			}

			switch op {
			case parenOpen:
				stack.Push(op)

			case parenClose:
				for stack.Count() > 0 && stack.MustPeek() != parenOpen {
					postfix = append(postfix, expressionNode{
						data: stack.MustPop(),
					})
				}
				_, ok := stack.Pop()
				if !ok {
					return nil, errors.New("stack was empty")
				}

			default:
				pushFunc(op)
			}

		default:
			return postfix, fmt.Errorf("invalid token `%v`", token.Str)
		}

	}

	for stack.Count() > 0 {
		postfix = append(postfix, expressionNode{
			data: stack.MustPop(),
		})
	}

	return postfix, nil
}

func Evaluate(expr Expression) (float64, error) {

	var operandStack stack.Stack[float64]

	for _, node := range expr {

		switch node.data.(type) {

		case operator:
			result, err := node.data.(operator).Apply(&operandStack)
			if err != nil {
				return 0, err
			}
			operandStack.Push(result)

		case constant:
			operandStack.Push(float64(node.data.(constant).value))

		case float64:
			operandStack.Push(node.data.(float64))

		default:
			return 0, fmt.Errorf("not implemented node type")
		}
	}

	result, ok := operandStack.Pop()
	if !ok {
		return 0, fmt.Errorf("failed to evaluate expression, stack was empty")
	}

	if operandStack.Count() > 0 {
		return 0, fmt.Errorf("failed to evaluate expression, stack was not empty: %v", operandStack)
	}

	return result, nil
}
