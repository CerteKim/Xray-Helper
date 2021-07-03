package main

import (
	"net"
	"os"
	"xrayd/app/model"
	"xrayd/app/router"
	"xrayd/common/config"
	"xrayd/common/log"

	"github.com/valyala/fasthttp"

	"github.com/takama/daemon"
)

type Service struct {
	daemon.Daemon
}

func (service *Service) Cmd() (string, error) {
	usage := "Usage: xrayd install | remove | start | stop | status"

	// if received any kind of command, do it
	if len(os.Args) > 1 {
		command := os.Args[1]
		switch command {
		case "install":
			return service.Install()
		case "remove":
			return service.Remove()
		case "start":
			return service.Start()
		case "stop":
			return service.Stop()
		case "status":
			return service.Status()
		default:
			return name + "\n" + description + "\n" + usage, nil
		}
	}

	initXrayD()

	xrayd := &XrayD{}
	return service.Run(xrayd)
}

func acceptConnection(listener net.Listener, listen chan<- net.Conn) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		listen <- conn
	}
}

func handleClient(client net.Conn) {
	if err := fasthttp.ServeConn(client, router.Router(&fasthttp.RequestCtx{}).Handler); err != nil {
		log.Errlog.Println(err)
	}
}

func initXrayD() {
	config.InitConfig()
	model.InitDB()
}
