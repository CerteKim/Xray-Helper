package xray

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
	"xrayd/common/log"

	"github.com/spf13/viper"
)

var exeCmd *exec.Cmd

func Start() string {
	if err := Stop(); err != nil {
		return fmt.Sprintln(err)
	}
	exeCmd = exec.Command(viper.GetString("xray.path"), "-confdir", viper.GetString("xray.confdir"))
	exeCmd.Env = append(os.Environ(), "XRAY_LOCATION_ASSET="+viper.GetString("xray.asset"))
	stdout, _ := exeCmd.StdoutPipe()
	if err := exeCmd.Start(); err != nil {
		return fmt.Sprintln(err)
	}
	r := bufio.NewReader(stdout)
	lines := new([]string)
	go readInfo(r, lines)
	status := make(chan struct{})
	go checkProc(exeCmd, status)
	stopper := time.NewTimer(time.Millisecond * 300)
	select {
	case <-stopper.C:
		return "success"
	case <-status:
		if err := Stop(); err != nil {
			return fmt.Sprintln(err)
		}
		return strings.Join(*lines, "\n")
	}
}

func Stop() error {
	if exeCmd != nil {
		err := exeCmd.Process.Kill()
		if err != nil {
			return err
		}
		exeCmd = nil
	}
	return nil
}

func readInfo(r *bufio.Reader, lines *[]string) {
	for i := 0; i < 20; i++ {
		line, _, _ := r.ReadLine()
		if len(string(line[:])) != 0 {
			*lines = append(*lines, string(line[:]))
		}
	}
}

// 检查进程状态
func checkProc(c *exec.Cmd, status chan struct{}) {
	if err := c.Wait(); err != nil {
		log.Errlog.Println(err)
	}
	status <- struct{}{}
}
