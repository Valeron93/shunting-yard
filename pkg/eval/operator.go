package eval

import (
	"math"

	"github.com/Valeron93/shunting-yard/pkg/stack"
)

type operator interface {
	Apply(operandStack *stack.Stack[float64]) (float64, error)
	Precedence() int32
}

var defaultOperatorMap = map[string]operator{
	"sin":   NewUnaryFunction("sin", math.Sin),
	"asin":  NewUnaryFunction("asin", math.Asin),
	"acos":  NewUnaryFunction("acos", math.Acos),
	"cos":   NewUnaryFunction("cos", math.Cos),
	"tan":   NewUnaryFunction("tan", math.Tan),
	"atan":  NewUnaryFunction("atan", math.Atan),
	"exp":   NewUnaryFunction("exp", math.Exp),
	"sqrt":  NewUnaryFunction("sqrt", math.Sqrt),
	"log":   NewUnaryFunction("log", math.Log),
	"floor": NewUnaryFunction("floor", math.Floor),
	"ceil":  NewUnaryFunction("ceil", math.Ceil),
	"abs":   NewUnaryFunction("abs", math.Abs),
	"(":     parenOpen,
	")":     parenClose,
	"+":     plus,
	"-":     minus,
	"*":     mul,
	"/":     div,
	"%":     mod,
	"mod":   mod,
	"^":     pow,
}
