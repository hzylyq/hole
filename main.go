package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/urfave/cli"

	"github.com/hzylyq/hole/entity"
)

var commends = []cli.Command{
	{
		Name:   "entity",
		Usage:  "input entity file",
		Action: entity.Entity,
	},
}



func main() {
	app := cli.NewApp()
	app.Usage = "a cli tool to generate code"
	app.Version = fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)
	app.Commands = commends

	// cli already print error messages
	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	log.Print("gen code success")
}
