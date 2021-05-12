package main

import (
	"log"
	"os"

	_ "xrayd/app/service"
)

const (
	name        = "xrayd"
	description = "An xray daemon"
)

var stdlog, errlog *log.Logger

func main() {
	srv := new(XrayD)
	status, err := srv.Cmd()
	if err != nil {
		errlog.Println(status, "\nError: ", err)
		os.Exit(1)
	}
	stdlog.Println(status)
}

func init() {
	stdlog = log.New(os.Stdout, "", log.Ldate|log.Ltime)
	errlog = log.New(os.Stderr, "", log.Ldate|log.Ltime)
}
