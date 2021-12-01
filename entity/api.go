package entity

import (
	"text/template"

	"github.com/urfave/cli"

	"github.com/hzylyq/hole/util"
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
	fileName := ctx.String("entity")
	if len(fileName) == 0 {
		panic("must input entity file")
	}

	fileByte, err := util.ReadFile(fileName)
	if err != nil {
		return err
	}

	template, err := template.New("cli").Parse(apiTemplate)
	if err != nil {
		return err
	}

	template.Parse(string(fileByte))

	return nil
}
