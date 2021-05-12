package tproxy

import "runtime"

type linuxTProxy struct{}

func makeITProxy() ITProxy {
	return &linuxTProxy{}
}

func (p *linuxTProxy) addRoute() error {
	return nil
}

func (p *linuxTProxy) delRoute() error {
	return nil
}

func (p *linuxTProxy) enableProxy() error {
	if runtime.GOOS == "android" {
		return nil
	}
	return nil
}

func (p *linuxTProxy) disableProxy() error {
	if runtime.GOOS == "android" {
		return nil
	}
	return nil
}
