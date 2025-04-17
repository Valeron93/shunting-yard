package tokenizer

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

const AllowedSimpleOperators = "()*-+/%^"

type Token struct {
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
	return fmt.Sprintf("%v", t.Data)
}

func Tokenize(input []rune) ([]Token, error) {
	tokens := make([]Token, 0, 10)

	i := 0
	for i < len(input) {
		c := input[i]

		if unicode.IsSpace(c) || c == ',' {
			i++
			continue
		}
		if unicode.IsNumber(c) || c == '.' {
			count, num, err := parseNumber(input[i:])
			if err != nil {
				return nil, err
			}

			tokens = append(tokens, Token{
				Data: num,
			})

			i += count
			continue
		}

		if strings.ContainsRune(AllowedSimpleOperators, c) {
			tokens = append(tokens, Token{
				Data: string(c),
			})
			i++
			continue
		}

		var f strings.Builder
		for _, s := range input[i:] {
			if strings.ContainsRune(AllowedSimpleOperators, s) {
				break
			}
			if unicode.IsSpace(s) || s == ',' {
				break
			}
			f.WriteRune(s)
		}

		data := f.String()
		if len(data) > 0 {
			tokens = append(tokens, Token{
				Data: f.String(),
			})
			i += len(data)
		}
	}

	return tokens, nil
}
