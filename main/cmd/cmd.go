package cmd

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"xrayd/common/config"

	"github.com/spf13/viper"
	"github.com/takama/daemon"
)

var stdlog, errlog *log.Logger

func init() {
	stdlog = log.New(os.Stdout, "", log.Ldate|log.Ltime)
	errlog = log.New(os.Stderr, "", log.Ldate|log.Ltime)
}

type Service struct {
	daemon.Daemon
}

func (srv *Service) Start() error {
	return nil
}

func (srv *Service) Stop() error {
	return nil
}

func (srv *Service) Status() error {
	return nil
}

func (srv *Service) Manage() (string, error) {

	usage := "xrayd version " + "0.0.1\n" + "Usage: xrayd start | stop | status"

	// if received any kind of command, do it
	if len(os.Args) > 1 {
		cmd := os.Args[1]
		switch cmd {
		case "start":
			config.LoadConfig()
			if err := srv.Start(); err != nil {
				return "Start failed", err
			}
		case "stop":
			if err := srv.Stop(); err != nil {
				return "Stop failed", err
			}
		case "status":
			if err := srv.Status(); err != nil {
				return "Failed to check status", err
			}
		default:
			return usage, nil
		}
	}

	// Do something, call your goroutines, etc

	// Set up channel on which to send signal notifications.
	// We must use a buffered channel or risk missing the signal
	// if we're not ready to receive when the signal is sent.
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)

	// Set up listener for defined host and port
	listener, err := net.Listen("tcp", ":"+viper.GetString("port"))
	if err != nil {
		return "Possibly was a problem with the port binding", err
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
				return "Daemon was interruped by system signal", nil
			}
			return "Daemon was killed", nil
		}
	}
}

// Accept a client connection and collect it in a channel
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
	for {
		buf := make([]byte, 4096)
		numbytes, err := client.Read(buf)
		if numbytes == 0 || err != nil {
			return
		}
		if _, err := client.Write(buf[:numbytes]); err != nil {
			log.Fatal(err)
		}
	}
}
