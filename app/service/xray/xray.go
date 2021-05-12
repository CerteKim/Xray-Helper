package xray

import (
	"bytes"
	"os"
	"os/exec"

	"github.com/spf13/viper"
)

type Xray interface {
	Start() error
	Stop() error
}

type xray struct {
}

var X xray

var Cmd = exec.Command(viper.GetString("xray.path"), "-confdir", viper.GetString("xray.confdir"))

func (x xray) Start() (string, string, error) {

	Cmd.Env = append(os.Environ(), "XRAY_LOCATION_ASSET="+viper.GetString("xray.asset"))

	var stdout, stderr bytes.Buffer
	Cmd.Stdout = &stdout
	Cmd.Stderr = &stderr

	err := Cmd.Start()
	if err != nil {
		return "", "", err
	}

	outStr, errStr := stdout.String(), stderr.String()

	return outStr, errStr, nil
}

func (x xray) Stop() error {
	if err := Cmd.Process.Kill(); err != nil {
		return err
	}
	return nil
}
