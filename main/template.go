package main

import "os"

func template() string {
	if _, err := os.Stat("/run/systemd/system"); err == nil {
		return systemDConfig
	}
	if _, err := os.Stat("/sbin/initctl"); err == nil {
		return upstatConfig
	}
	return systemVConfig
}
