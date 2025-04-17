package eval

import (
	"fmt"

	"github.com/Valeron93/shunting-yard/pkg/stack"
)

type UnaryFunction struct {
	f    func(float64) float64
	name string
}

func (f UnaryFunction) Apply(operandStack *stack.Stack[float64]) (float64, error) {
	operand, ok := operandStack.Pop()

	if !ok {
		return 0, fmt.Errorf("function `%v` expects one argument", f.name)
	}
	return f.f(operand), nil
}

func (f UnaryFunction) Precedence() int32 {
	return 6
}

func (f UnaryFunction) String() string {
	return f.name
}

func NewUnaryFunction(name string, f func(float64) float64) UnaryFunction {
	return UnaryFunction{
		f:    f,
		name: name,
	}
}
