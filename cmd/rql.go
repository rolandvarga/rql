package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github/rolandvarga/rql/internal/parser"
)

const (
	PROMPT              = "rql> "
	STATEMENT_SEPERATOR = ";"

	BUILTIN_PREFIX = "."
	BUILTIN_EXIT   = ".exit"
	BUILTIN_HELP   = ".help"
	BUILTIN_TABLES = ".tables"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("rql> ")

		var input string
		for scanner.Scan() {
			input += scanner.Text()

			if strings.Contains(input, STATEMENT_SEPERATOR) {
				if !strings.HasSuffix(input, STATEMENT_SEPERATOR) {
					fmt.Println("Error: invalid input")
					break
				}
				input = strings.TrimSuffix(input, STATEMENT_SEPERATOR)
				break
			}
		}

		statement, err := parser.ParseInput(input)
		if err != nil {
			fmt.Println(err)
			continue
		}

		statement.Execute()
	}
}
