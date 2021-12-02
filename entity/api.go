package entity

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"

	"github.com/urfave/cli"
)

type Field struct {
	Name  string
	Value string
}

type entityStruct struct {
	TableName string
	PackName  string
	Fields    []Field
	hasDelete bool
}

const apiTemplate = `
type Request {
	Id int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
`

func Entity(ctx *cli.Context) error {
	log.Println("run here")
	fileName := ctx.String("o")
	if len(fileName) == 0 {
		panic("must input entity file")
	}

	ent, err := ReadEntityFromFile(fileName)
	if err != nil {
		return err
	}

	log.Print(ent)
	return nil
}

func ReadEntityFromFile(fileName string) (*entityStruct, error) {
	f, err := parser.ParseFile(token.NewFileSet(), fileName, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	ent := new(entityStruct)

	ent.PackName = f.Name.Name
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

		ent.TableName = structName

		for _, field := range s.Fields.List {
			ent.Fields = append(ent.Fields, Field{
				Name:  field.Names[0].Name,
				Value: field.T,
			})
			fmt.Printf("Field: %s\n", field.Names[0].Name)
			fmt.Printf("Tag:   %s\n", field.Tag.Value)
		}

		return false
	})

	return ent, nil
}
