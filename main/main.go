package main

import (
	"os"
	"runtime"
	"xrayd/common/log"

	"github.com/takama/daemon"
)

const (
	name        = "xrayd"
	description = "An xray daemon"
)

var dependencies = []string{}

func main() {
	daemonKind := daemon.SystemDaemon
	if runtime.GOOS == "darwin" {
		daemonKind = daemon.UserAgent
	}
	srv, err := daemon.New(name, description, daemonKind, dependencies...)
	if err != nil {
		log.Errlog.Println("Error: ", err)
		os.Exit(1)
	}
	if runtime.GOOS == "linux" {
		if err := srv.SetTemplate(template()); err != nil {
			log.Errlog.Println("Error: ", err)
			os.Exit(1)
		}
	}
	service := &Service{srv}
	status, err := service.Cmd()
	if err != nil {
		log.Errlog.Println(status, "\nError: ", err)
		os.Exit(1)
	}
	log.Stdlog.Println(status)
}
