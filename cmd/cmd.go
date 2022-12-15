package cmd

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

const (
	Insert = iota + 1
	Select
	Update
	Delete
)

type Statement struct {
	Type int
}

func (s *Statement) Execute() {}

func parseInput(statement string) (string, error) {
	return "", inputParseError
	return "", nil
}
