package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

var (
	inputParseError = errors.New("input cannot be parsed as a valid RQL statement")
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("rql> ")
		input, err := reader.ReadString(';')
		if err != nil {
			fmt.Printf("internal error: %v", err)
		}

		statement, err := parseInput(input)
		if err != nil {
			fmt.Println(err)
		}

		statement.Execute()
	}
}

type StatementType int

const (
	Insert StatementType = iota + 1
	Select
	Update
	Delete
)

type Statement struct {
	Type StatementType
}

func (s *Statement) Execute() {}

func parseInput(input string) (Statement, error) {
	if false {
		return Statement{}, inputParseError
	}
	return Statement{}, nil
}
