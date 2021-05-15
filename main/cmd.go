package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"xrayd/app/model"
	"xrayd/app/router"
	"xrayd/common/config"

	"github.com/spf13/viper"
	"github.com/takama/daemon"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
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
			initX()
			h2s := &http2.Server{}
			if err := http.ListenAndServe(":"+viper.GetString("xrayd.port"), h2c.NewHandler(router.NewRouter(), h2s)); err != nil {
				return fmt.Sprintln(err), nil
			}
		default:
			return name + "\n" + description + "\n" + usage, nil
		}
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	for {
		select {
		case killSignal := <-interrupt:
			stdlog.Println("Stopped")
			if killSignal == os.Interrupt {
				return "Daemon was interruped by system signal", nil
			}
			return "Daemon was killed", nil
		}
	}
}

func initX() {
	config.InitConfig()
	model.InitDB()
}
