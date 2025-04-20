package tokenizer

import (
	"fmt"
	"slices"
	"strings"
	"unicode"
)

var allowedOperators = []string{
	"(",
	")",
	"+",
	"-",
	"*",
	"/",
	"mod",
	"%",
	"^",
	"**",
}

type TokenType int32
type OperatorType int32

const (
	Invalid    TokenType = iota
	Operator   TokenType = iota
	Identifier TokenType = iota
	Number     TokenType = iota
)

type Token struct {
	Type TokenType
	Str  string
}

type Tokenizer struct {
	input  []rune
	cursor int
}

func (t Token) String() string {
	return fmt.Sprintf("%v", t.Str)

}

func (t *Tokenizer) canAdvance() bool {
	return t.cursor < len(t.input)
}

func (t *Tokenizer) getRune() rune {
	if t.canAdvance() {
		return t.input[t.cursor]
	}
	return 0
}

func (t *Tokenizer) skipWhitespacesAndCommas() {
	for unicode.IsSpace(t.getRune()) || t.getRune() == ',' {
		t.cursor++
	}
}

func (t *Tokenizer) getInvalidToken() Token {
	start := t.cursor
	for t.canAdvance() && !unicode.IsSpace(t.getRune()) {
		t.cursor++
	}

	return Token{
		Type: Invalid,
		Str:  t.getSubstringToCursor(start),
	}
}

func Tokenize(input string) []Token {
	t := Tokenizer{
		input:  []rune(input),
		cursor: 0,
	}
	var tokens []Token
	for {
		token, ok := t.NextToken()
		if !ok {
			break
		}
		tokens = append(tokens, token)
	}
	return tokens
}

func (t *Tokenizer) NextToken() (Token, bool) {

	t.skipWhitespacesAndCommas()
	if !t.canAdvance() {
		return Token{}, false
	}
	r := t.getRune()
	if canBeNumber(r) {
		return t.getNumber(), true
	}
	if canBeOperator(string(r)) {
		op, ok := t.getOperator()
		if ok {
			return op, true
		}
	}
	if canBeIdentifier(r) {
		return t.getIdentifier(), true
	}
	return t.getInvalidToken(), true
}

func (t *Tokenizer) getOperator() (Token, bool) {
	start := t.cursor

	for t.canAdvance() {
		t.cursor++
		if !canBeOperator(t.getSubstringToCursor(start)) {
			t.cursor--
			break
		}
	}
	s := t.getSubstringToCursor(start)
	if !slices.Contains(allowedOperators, s) {
		t.cursor = start
		return Token{}, false
	}

	return Token{
		Type: Operator,
		Str:  t.getSubstringToCursor(start),
	}, true
}

func (t *Tokenizer) getIdentifier() Token {
	start := t.cursor

	for t.canAdvance() {
		r := t.getRune()
		if r == '_' || unicode.IsDigit(r) || unicode.IsLetter(r) {
			t.cursor++
		} else {
			break
		}
	}

	return Token{
		Type: Identifier,
		Str:  t.getSubstringToCursor(start),
	}
}

func (t *Tokenizer) getNumber() Token {
	start := t.cursor

	scientific := false
	sign := false
	dot := false

	for t.canAdvance() {
		r := t.getRune()
		if unicode.IsDigit(r) {

		} else if r == '.' {
			if dot {
				break
			}
			dot = true
		} else if r == 'e' {
			if scientific {
				break
			}
			scientific = true
		} else if r == '+' || r == '-' {
			if !scientific || sign {
				break
			}
			sign = true
		} else {
			break
		}
		t.cursor++
	}

	s := t.getSubstringToCursor(start)
	return Token{
		Type: Number,
		Str:  s,
	}
}

func (t *Tokenizer) getSubstringToCursor(start int) string {
	return string(t.input[start:t.cursor])
}

func canBeNumber(r rune) bool {
	return r == '.' || unicode.IsDigit(r)
}

func canBeIdentifier(r rune) bool {
	return r == '_' || unicode.IsLetter(r)
}

// TODO: we can try using Trie to optimize this
func canBeOperator(token string) bool {
	for _, allowed := range allowedOperators {
		if strings.HasPrefix(allowed, token) {
			return true
		}
	}
	return false

}
