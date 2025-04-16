package operator

import (
	"fmt"
	"slices"
)

type Operator rune

const AllowedOperators = "()*-+/%^"

const (
	ParenOpen  Operator = '('
	ParenClose Operator = ')'
	Mul        Operator = '*'
	Minus      Operator = '-'
	Plus       Operator = '+'
	Div        Operator = '/'
	Mod        Operator = '%'
	Pow        Operator = '^'
)

func (op Operator) Precedence() int32 {
	switch op {
	case ParenOpen, ParenClose:
		return 1
	case Minus, Plus:
		return 2
	case Mul, Mod, Div:
		return 3
	case Pow:
		return 4
	default:
		return -9999
	}
}

func (op Operator) Apply(operand1, operand2 float64) (result float64, err error) {

	switch op {
	case Minus:
		result = operand1 - operand2
		return
	
	case Plus:
		result = operand1 + operand2
		return
	case Mul:
		result = operand1 * operand2
		return
	
	case Div:
		result = operand1 / operand2
		return
	
	default:
		err = fmt.Errorf("not implemented")
		return
	}
}

func (op Operator) String() string {
	return fmt.Sprintf("%v", rune(op))
}

func FromRune(op rune) (Operator, error) {

	if slices.Contains([]rune(AllowedOperators), op) {
		return Operator(op), nil
	}

	return 0, fmt.Errorf("unknown error: `%c`", op)
}


