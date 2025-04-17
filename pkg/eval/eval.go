package eval

import (
	"errors"
	"fmt"

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

func opStringToOperator(op string) (operator, error) {

	if f, ok := defaultOperatorMap[op]; ok {
		return f, nil
	}

	return nil, fmt.Errorf("unknown operator/function: %v", op)
}

func TokensToExpression(tokens []tokenizer.Token) (Expression, error) {
	postfix := make(Expression, 0, 10)
	var stack stack.Stack[operator]

	for _, token := range tokens {

		switch token.Data.(type) {

		case float64:
			postfix = append(postfix, expressionNode{
				data: token.Data,
			})

		case string:

			if num, ok := defaultConstantMap[token.Data.(string)]; ok {
				postfix = append(postfix, expressionNode{
					data: num,
				})
				continue
			}

			op, err := opStringToOperator(token.Data.(string))
			if err != nil {
				return nil, err
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
				for stack.Count() > 0 && op.Precedence() <= stack.MustPeek().Precedence() {
					postfix = append(postfix, expressionNode{
						data: stack.MustPop(),
					})
				}
				stack.Push(op)
			}
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
