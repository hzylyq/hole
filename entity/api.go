package entity

import (
	"log"

	"github.com/urfave/cli"

	"github.com/hzylyq/hole/util"
)

type entityStruct struct {
	SchemaName string
	Field      []string
}

func Entity(ctx *cli.Context) error {
	fileName := ctx.String("entity")
	if len(fileName) == 0 {
		panic("must input entity file")
	}

	// todo read file gen entity
	fileByte, err := util.ReadFile(fileName)
	if err != nil {
		return err
	}

	log.Print(fileByte)

	return nil
}
