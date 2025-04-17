# Shunting Yard Algorithm

## Description

A simple shunting yard algorithm implementation in Go. This algorithm is used to parse and evaluate mathematical expressions.
Program can be used a library, or terminal UI calculator application.

Implemented features:
- Number parsing
- Basic arithmetic operators
- Basic mathematical functions (sin, cos, sqrt, floor, etc...)
- Basic mathematical constants (pi, phi, e, etc...)

## TODO:
- [x] Constant expressions (e.g. Pi, Euler's, Phi)
- [ ] More meaningful error messages
- [ ] Unit tests
- [ ] Add history to TUI

## Building and running the program

### Run terminal UI
```shell
$ go run ./cmd/ui
```

### Run examples
```shell
$ go run ./cmd/examples
```

## References:
- [Shunting Yard Algorithm (Wikipedia)](https://en.wikipedia.org/wiki/Shunting_yard_algorithm)
- [bubbletea (Go TUI Library)](https://github.com/charmbracelet/bubbletea)

## Examples demo: 

```shell
$ go run ./cmd/examples

------ EXAMPLE 1 ------
example: "1+1"
tokenizer output: [1 + 1]
parsed expression: [1 1 +]
result: 2
------------------------

------ EXAMPLE 2 ------
example: "-2 + 1"
tokenizer output: [- 2 + 1]
parsed expression: [2 - 1 +]
result: -1
------------------------

------ EXAMPLE 3 ------
example: "sin(1000)^2 + cos(1000)^2 + 10^2 + 3 mod 2"
tokenizer output: [sin ( 1000 ) ^ 2 + cos ( 1000 ) ^ 2 + 10 ^ 2 + 3 mod 2]
parsed expression: [1000 sin 2 ^ 1000 cos 2 ^ + 10 2 ^ + 3 2 % +]
result: 102
------------------------

------ EXAMPLE 4 ------
example: "2*2*2"
tokenizer output: [2 * 2 * 2]
parsed expression: [2 2 * 2 *]
result: 8
------------------------

------ EXAMPLE 5 ------
example: "3 + 4 * 2 / (1 - 5)^2"
tokenizer output: [3 + 4 * 2 / ( 1 - 5 ) ^ 2]
parsed expression: [3 4 2 * 1 5 - 2 ^ / +]
result: 3.5
------------------------

------ EXAMPLE 6 ------
example: "sqrt(16) + log(100)"
tokenizer output: [sqrt ( 16 ) + log ( 100 )]
parsed expression: [16 sqrt 100 log +]
result: 8.605170185988092
------------------------

------ EXAMPLE 7 ------
example: "tan(45) + atan(1)"
tokenizer output: [tan ( 45 ) + atan ( 1 )]
parsed expression: [45 tan 1 atan +]
result: 2.4051733539413096
------------------------

------ EXAMPLE 8 ------
example: "abs(-42) + 7 mod 5"
tokenizer output: [abs ( - 42 ) + 7 mod 5]
parsed expression: [42 - abs 7 5 % +]
result: 44
------------------------

------ EXAMPLE 9 ------
example: "5 + 2^3"
tokenizer output: [5 + 2 ^ 3]
parsed expression: [5 2 3 ^ +]
result: 13
------------------------

------ EXAMPLE 10 ------
example: "exp(1)^2 - log(7)"
tokenizer output: [exp ( 1 ) ^ 2 - log ( 7 )]
parsed expression: [1 exp 2 ^ 7 log -]
result: 5.443145949875336
------------------------

------ EXAMPLE 11 ------
example: "floor(9.9) + ceil(1.1)"
tokenizer output: [floor ( 9.9 ) + ceil ( 1.1 )]
parsed expression: [9.9 floor 1.1 ceil +]
result: 11
------------------------

------ EXAMPLE 12 ------
example: "exp(log(100))"
tokenizer output: [exp ( log ( 100 ) )]
parsed expression: [100 log exp]
result: 100.00000000000004
------------------------

------ EXAMPLE 13 ------
example: "1 + sin 2"
tokenizer output: [1 + sin 2]
parsed expression: [1 2 sin +]
result: 1.9092974268256815
------------------------

------ EXAMPLE 14 ------
example: "pi + 1"
tokenizer output: [pi + 1]
parsed expression: [pi 1 +]
result: 4.141592653589793
------------------------

------ EXAMPLE 15 ------
example: "exp(2.2)"
tokenizer output: [exp ( 2.2 )]
parsed expression: [2.2 exp]
result: 9.025013499434122
------------------------

------ EXAMPLE 16 ------
example: "e^2.2"
tokenizer output: [e ^ 2.2]
parsed expression: [e 2.2 ^]
result: 9.025013499434122
------------------------

------ EXAMPLE 17 ------
example: "phi"
tokenizer output: [phi]
parsed expression: [phi]
result: 1.618033988749895
------------------------

------ EXAMPLE 18 ------
example: "(1+sqrt(5))/2"
tokenizer output: [( 1 + sqrt ( 5 ) ) / 2]
parsed expression: [1 5 sqrt + 2 /]
result: 1.618033988749895
------------------------
```