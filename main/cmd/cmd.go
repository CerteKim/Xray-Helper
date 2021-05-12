package cmd

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"xrayd/app/service"
	"xrayd/common/config"

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
	return service.X.Start()
}

func (srv *Service) Stop() error {
	return service.X.Stop()
}

func (srv *Service) Status() (int, error) {
	return service.Pid.Pid, nil
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
			if _, err := srv.Status(); err != nil {
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
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	killSignal := <-interrupt
	stdlog.Println("Got signal:", killSignal)
	return "Service exited", nil
}
