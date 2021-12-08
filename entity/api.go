package entity

import (
	"log"

	"github.com/urfave/cli"

	"github.com/hzylyq/hole/core/entity"
)

func Entity(ctx *cli.Context) error {
	log.Println("run here")
	fileName := ctx.String("o")
	if len(fileName) == 0 {
		panic("must input entity file")
	}

	ent, err := entity.NewEntityFromFile(fileName)
	if err != nil {
		return err
	}

	log.Print(ent)
	return nil
}
