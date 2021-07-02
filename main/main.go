package main

import (
	"log"
	"os"

	"xrayd/app/model"
	_ "xrayd/app/service"
	"xrayd/common/config"

	"github.com/takama/daemon"
)

const (
	name        = "xrayd"
	description = "An xray daemon"
)

var stdlog, errlog *log.Logger

func main() {
	srv, err := daemon.New(name, description, daemon.SystemDaemon)
	if err != nil {
		errlog.Println("Error: ", err)
		os.Exit(1)
	}
	service := &XrayD{srv}
	status, err := service.Cmd()
	if err != nil {
		errlog.Println(status, "\nError: ", err)
		os.Exit(1)
	}
	stdlog.Println(status)
}

func init() {
	stdlog = log.New(os.Stdout, "", log.Ldate|log.Ltime)
	errlog = log.New(os.Stderr, "", log.Ldate|log.Ltime)
	config.InitConfig()
	model.InitDB()
}
