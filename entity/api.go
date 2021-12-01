package entity

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"

	"github.com/urfave/cli"
)

type entityStruct struct {
	SchemaName string
	PackName   string
	Field      []string
	hasDelete  bool
}

const apiTemplate = `
type Request {
	Id int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
`

var ent entityStruct

func Entity(ctx *cli.Context) error {
	log.Println("run here")
	fileName := ctx.String("o")
	if len(fileName) == 0 {
		panic("must input entity file")
	}

	f, err := parser.ParseFile(token.NewFileSet(), fileName, nil, parser.ParseComments)
	if err != nil {
		return err
	}

	packageName := f.Name.Name

	ent.PackName = packageName
	var structName string

	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.TypeSpec:
			if x.Type == nil {
				return true
			}
			structName = x.Name.Name
		}

		s, ok := n.(*ast.StructType)
		if !ok {
			return true
		}

		log.Println(structName)

		for _, field := range s.Fields.List {

			fmt.Printf("Field: %s\n", field.Names[0].Name)
			fmt.Printf("Tag:   %s\n", field.Tag.Value)
		}

		return false
	})

	return nil
}
