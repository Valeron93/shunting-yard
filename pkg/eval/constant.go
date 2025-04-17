package eval

import "math"

var defaultConstantMap = map[string]constant{
	"pi": {
		name:  "pi",
		value: math.Pi,
	},
	"e": {
		name:  "e",
		value: math.E,
	},
	"phi": {
		name:  "phi",
		value: math.Phi,
	},
}

type constant struct {
	name  string
	value float64
}

func (c constant) String() string {
	return c.name
}
