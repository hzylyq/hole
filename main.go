package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"

	"github.com/hzylyq/hole/entity"
)

var commends = []cli.Command{
	{
		Name:   "entity",
		Usage:  "input entity file",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "o",
				Usage: "the output api file",
			},
		},
		Action: entity.Entity,
	},
}

func main() {
	app := cli.NewApp()
	app.Name = "hole"
	app.Usage = "a cli tool to generate code"
	app.Commands = commends

	// cli already print error messages
	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	log.Print("gen code success")
}
