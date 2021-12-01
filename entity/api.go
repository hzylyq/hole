package entity

import (
	"go/parser"
	"go/token"
	"log"

	"github.com/urfave/cli"
)

type entityStruct struct {
	SchemaName string
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

	log.Println(packageName)

	// 找到这个文件的package path

	// template, err := template.New("cli").Parse(apiTemplate)
	// if err != nil {
	// 	return err
	// }
	//
	// template.Parse(string(fileByte))

	return nil
}
