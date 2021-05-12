package main

import (
	"fmt"
	"net/http"
	"os"
	"xrayd/app/model"
	"xrayd/app/router"
	"xrayd/common/config"

	"github.com/spf13/viper"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type XrayD struct{}

func (X XrayD) Cmd() (string, error) {
	usage := "Usage: xrayd run"

	// if received any kind of command, do it
	if len(os.Args) > 1 {
		command := os.Args[1]
		switch command {
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
	return name + "\n" + description + "\n" + usage, nil
}

func initX() {
	config.InitConfig()
	model.InitDB()
}
