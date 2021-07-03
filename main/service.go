package main

import (
	"net"
	"os"
	"os/signal"
	"syscall"
	"xrayd/common/log"

	"github.com/spf13/viper"
)

type XrayD struct {
	listen chan net.Conn
}

func (X *XrayD) Start() {
	listener, err := net.Listen("tcp", ":"+viper.GetString("xrayd.port"))
	if err != nil {
		log.Errlog.Println("Possibly was a problem with the port binding", err)
	} else {
		log.Stdlog.Println("Listening on ", listener.Addr())
	}

	X.listen = make(chan net.Conn, 100)
	go acceptConnection(listener, X.listen)

	go func() {
		for {
			conn, ok := <-X.listen
			if !ok {
				log.Stdlog.Println("Closing connections")
				listener.Close()
				return
			}
			go handleClient(conn)
		}
	}()
}

func (X *XrayD) Stop() {
	close(X.listen)
}

func (X *XrayD) Run() {
	X.Start()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

loop:
	for killSignal := range interrupt {
		log.Stdlog.Println("Got signal:", killSignal)
		if killSignal == os.Interrupt {
			log.Stdlog.Println("Daemon was interrupted by system signal")
			break loop
		}
		log.Stdlog.Println("Daemon was killed")
	}

	X.Stop()
}
