package eval

import (
	"fmt"
	"math"

	"github.com/Valeron93/shunting-yard/pkg/stack"
)

type arithmetic rune

const (
	parenOpen  arithmetic = '('
	parenClose arithmetic = ')'
	mul        arithmetic = '*'
	minus      arithmetic = '-'
	plus       arithmetic = '+'
	div        arithmetic = '/'
	mod        arithmetic = '%'
	pow        arithmetic = '^'
)

func (op arithmetic) Precedence() int32 {
	switch op {
	case parenOpen, parenClose:
		return 1
	case minus, plus:
		return 2
	case mul, mod, div:
		return 3
	case pow:
		return 4
	default:
		return -9999
	}
}

func (op arithmetic) Apply(operandStack *stack.Stack[float64]) (result float64, err error) {

	if op == minus && operandStack.Count() == 1 {
		return -operandStack.MustPop(), nil
	}

	if operandStack.Count() < 2 {
		return 0, fmt.Errorf("operator %v expects 2 arguments, got %v", op, operandStack.Count())
	}

	operand2 := operandStack.MustPop()
	operand1 := operandStack.MustPop()

	switch op {
	case minus:
		result = operand1 - operand2

	case plus:
		result = operand1 + operand2

	case mul:
		result = operand1 * operand2

	case div:
		result = operand1 / operand2

	case pow:
		result = math.Pow(operand1, operand2)

	case mod:
		result = math.Mod(operand1, operand2)

	default:
		err = fmt.Errorf("operator %c is unknown or not implemented", op)
	}
	return
}

func (op arithmetic) String() string {
	return fmt.Sprintf("%c", rune(op))
}
