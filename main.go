package main

import (
	"flag"
	"log"
	"os"
)

const defaultEntity = ""
const defaultServiceDir = "./"

func main() {
	entityFile := flag.String("entity", defaultEntity, "entity go file")
	serviceDir := flag.String("service", defaultServiceDir, "service dir")
	flag.Parse()

	if entityFile == nil {
		log.Fatalln("must input entity file.")
	}

	// 读取entityFile内容 解析成schema
	curWd, err := os.Getwd()
	if err != nil {
		log.Fatalln("can't get current working directory: ", err)
	}

	file, err := os.Open(curWd + *entityFile)
	if err != nil {
		log.Fatalln("filed to read entity file")
	}

	// 解析file

	log.Print(entityFile)
	log.Print(serviceDir)

	print("hello hole")
}
