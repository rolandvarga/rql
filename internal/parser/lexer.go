package parser

import (
	"errors"
	"github/rolandvarga/rql/internal/model"
	"strings"
)

var (
	inputParseError = errors.New("input cannot be parsed as a valid RQL statement")
)

var (
	LEXER_SEPARATOR = map[rune]bool{
		' ':  true,
		'.':  true,
		',':  true,
		'\n': true,
		'\t': true,
	}

	LEXER_SKIP = map[string]bool{
		"":   true,
		",":  true,
		";":  true,
		"\n": true,
		"\t": true,
	}
)

func LexInput(input string) []string {
	var tokens []string

	input = strings.ToLower(input)
	parts := strings.FieldsFunc(input, func(c rune) bool {
		return LEXER_SEPARATOR[c]
	})

	for _, part := range parts {
		if LEXER_SKIP[part] {
			continue
		}
		tokens = append(tokens, part)
	}
	return tokens
}

func ParseInput(input string) (model.Statement, error) {
	tokens := LexInput(input)

	if len(tokens) == 0 {
		return model.Statement{}, inputParseError
	}

	return model.NewStatement(tokens)
}
