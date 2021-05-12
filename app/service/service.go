package service

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Xray interface {
	Start() error
	Stop() error
}

type xray struct {
}

var Pid *os.Process

var X xray

func (x xray) Start() error {
	procAttr := &os.ProcAttr{
		Env: {
			"XRAY_LOCATION_ASSET=" + viper.GetString("xray.asset"),
		},
		Files: []*os.File{
			os.Stdin,
			os.Stdout,
			os.Stderr,
		},
	}

	var err error
	Pid, err = os.StartProcess(viper.GetString("xray.path"), []string{"xray", "-confdir", viper.GetString("xray.confdir")}, procAttr)
	if err != nil {
		log.Printf("Error %v starting process!", err)
		os.Exit(1)
	}
	log.Printf("The process id is %v", Pid)
}

func (x xray) Stop() error {

}
