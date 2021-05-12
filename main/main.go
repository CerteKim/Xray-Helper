package main

import (
	"fmt"
	"log"
	"os"

	_ "xrayd/common/config"

	"xrayd/main/cmd"

	"github.com/takama/daemon"
)

const (
	name        = "xrayd"
	description = "An xray daemon"
)

var dependencies = []string{"dummy.service"}

func main() {
	srv, err := daemon.New(name, description, daemon.SystemDaemon, dependencies...)
	if err != nil {
		log.Fatal("Error: ", err)
		os.Exit(1)
	}
	service := &cmd.Service{srv}
	status, err := service.Manage()
	if err != nil {
		log.Fatal(status, "\nError: ", err)
		os.Exit(1)
	}
	fmt.Println(status)
}
