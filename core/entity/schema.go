package entity

import "go/ast"

type field struct {
	Name  string
	Value ast.Expr
}

type schema struct {
	TableName string
	PackName  string
	Fields    []field
	hasDelete bool
}
