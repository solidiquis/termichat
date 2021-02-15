package main

import (
	"log"
	"os"
)

const (
	DEFAULT_PORT = 8080
)

var (
	infoLog *log.Logger
	errLog  *log.Logger
)

func init() {
	infoLog = log.New(
		os.Stdout,
		"\x1b[31mINFO\x1b[0m:\t",
		log.Ldate|log.Ltime,
	)

	errLog = log.New(
		os.Stderr,
		"\x1b[36mERROR\x1b[0m:\t",
		log.Ldate|log.Ltime|log.Lshortfile,
	)
}
