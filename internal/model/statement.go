package model

import (
	"errors"
	"fmt"
)

var (
	invalidStatementError = errors.New("invalid statement")
)

type StatementType int

const (
	Insert StatementType = iota + 1
	Select
	Update
	Delete
)

func (st StatementType) String() string {
	switch st {
	case Insert:
		return "Insert"
	case Select:
		return "Select"
	case Update:
		return "Update"
	case Delete:
		return "Delete"
	}
	return "Unknown"
}

type Statement struct {
	Type   StatementType
	Tokens []string
}

func NewStatement(tokens []string) (Statement, error) {
	statementType, err := getStatementTypeFrom(tokens[0])
	if err != nil {
		return Statement{}, err
	}
	return Statement{Type: statementType, Tokens: tokens}, nil
}

func (s *Statement) Execute() {
	fmt.Printf("statement type: %v\n", s.Type)
}

func getStatementTypeFrom(token string) (StatementType, error) {
	switch token {
	case "insert":
		return Insert, nil
	case "select":
		return Select, nil
	case "update":
		return Update, nil
	case "delete":
		return Delete, nil
	}
	return 0, invalidStatementError
}
