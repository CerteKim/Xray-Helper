package main

import (
	"os"
	"xrayd/common/log"

	"github.com/takama/daemon"
)

const (
	name        = "xrayd"
	description = "An xray daemon"
)

func main() {
	srv, err := daemon.New(name, description, daemon.SystemDaemon)
	if err != nil {
		log.Errlog.Println("Error: ", err)
		os.Exit(1)
	}
	srv.SetTemplate(template())
	service := &XrayD{srv}
	status, err := service.Cmd()
	if err != nil {
		log.Errlog.Println(status, "\nError: ", err)
		os.Exit(1)
	}
	log.Stdlog.Println(status)
}
