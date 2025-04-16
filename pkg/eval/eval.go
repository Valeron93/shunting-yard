package eval

import (
	"errors"

	"github.com/Valeron93/shunting-yard/pkg/operator"
	"github.com/Valeron93/shunting-yard/pkg/stack"
	"github.com/Valeron93/shunting-yard/pkg/tokenizer"
)


func Evaluate(tokens []tokenizer.Token) (float64, error) {
	
	var stack stack.Stack[float64]

	for _, token := range tokens {
		if token.Type == tokenizer.Number {
			stack.Push(token.Data.(float64))
			continue
		}

		if token.Type == tokenizer.Op {
			operand2, ok := stack.Pop()
			if !ok {
				return 0, errors.New("failed to get second operand, stack is malformed")
			}
			operand1, ok := stack.Pop()
			if !ok {
				return 0, errors.New("failed to get first operand, stack is malformed")
			}

			op := token.Data.(operator.Operator)
			result, err := op.Apply(operand1, operand2)
			if err != nil {
				return 0, err
			}
			stack.Push(result)
		}
	}
	result, ok := stack.Pop()
	if !ok {
		return 0, errors.New("failed to get result, stack is malformed")
	}
	if stack.Count() > 0 {
		return 0, errors.New("stack was not empty")
	}

	return result, nil

}