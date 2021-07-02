package main

import (
	"net"
	"os"
	"syscall"
	"xrayd/app/router"

	"github.com/valyala/fasthttp"

	"os/signal"

	"github.com/spf13/viper"
	"github.com/takama/daemon"
)

type XrayD struct {
	daemon.Daemon
}

func (X XrayD) Cmd() (string, error) {
	usage := "Usage: xrayd install | remove | start | stop | status"

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
		default:
			return name + "\n" + description + "\n" + usage, nil
		}
	}

	// Do something, call your goroutines, etc

	// Set up channel on which to send signal notifications.
	// We must use a buffered channel or risk missing the signal
	// if we're not ready to receive when the signal is sent.
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	// Set up listener for defined host and port
	listener, err := net.Listen("tcp", ":"+viper.GetString("xrayd.port"))
	if err != nil {
		return "Possibly was a problem with the port binding", err
	} else {
		stdlog.Println("Listening on ", listener.Addr())
	}

	// set up channel on which to send accepted connections
	listen := make(chan net.Conn, 100)
	go acceptConnection(listener, listen)

	// loop work cycle with accept connections or interrupt
	// by system signal
	for {
		select {
		case conn := <-listen:
			go handleClient(conn)
		case killSignal := <-interrupt:
			stdlog.Println("Got signal:", killSignal)
			stdlog.Println("Stoping listening on ", listener.Addr())
			listener.Close()
			if killSignal == os.Interrupt {
				return "Daemon was interrupted by system signal", nil
			}
			return "Daemon was killed", nil
		}
	}
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
	if err := fasthttp.ServeConn(client, func(ctx *fasthttp.RequestCtx) {
		router.Router(ctx)
	}); err != nil {
		os.Exit(1)
	}
}
