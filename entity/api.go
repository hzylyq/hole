package entity

import (
	"log"
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
  Name string ` + "`" + `path:"name,options=you|me"` + "`" + ` 
}

type Response {
  Message string ` + "`" + `json:"message"` + "`" + `
}

service {{.name}}-api {
  @handler {{.handler}}Handler
  get /from/:name(Request) returns (Response)
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

	log.Print(template)
	log.Print(fileByte)

	return nil
}
