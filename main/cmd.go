package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"xrayd/app/model"
	"xrayd/app/router"
	"xrayd/common/config"

	"github.com/spf13/viper"
	"github.com/takama/daemon"
)

type XrayD struct {
	daemon.Daemon
}

func (X XrayD) Cmd() (string, error) {
	usage := "Usage: xrayd install | remove | start | stop | status | run"

	// if received any kind of command, do it
	if len(os.Args) > 1 {
		command := os.Args[1]
		switch command {
		case "install":
			return X.Install()
		case "remove":
			return X.Remove()
		case "start":
			return X.Start()
		case "stop":
			return X.Stop()
		case "status":
			return X.Status()
		case "run":
			return X.Run()
		default:
			return name + "\n" + description + "\n" + usage, nil
		}
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		<-c
		log.Println("Exit")
		os.Exit(1)
	}()

	ret, err := X.Run()
	if err != nil {
		return "Error starting xrayd", err
	} else {
		return ret, err
	}
}

func initX() {
	config.InitConfig()
	model.InitDB()
}

func (X XrayD) Run() (string, error) {
	initX()
	err := http.ListenAndServe(":"+viper.GetString("xrayd.port"), router.NewRouter())
	return fmt.Sprintln(err), nil
}
