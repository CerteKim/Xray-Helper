package tproxy

import "runtime"

type linuxTProxy struct{}

func makeITProxy() ITProxy {
	return &linuxTProxy{}
}

func (t *linuxTProxy) AddRoute() error {
	//add_route()
	return nil
}

func (t *linuxTProxy) DelRoute() error {
	//del_route()
	return nil
}

func (t *linuxTProxy) EnableProxy() error {
	//create_mangle_iptables()
	if err := enableProxy(); err != nil {
		return err
	}
	if runtime.GOOS == "android" {
		Proxy.ApplyProxy()
		// ${iptables} -t mangle -A OUTPUT -j PROXY
		return nil
	}
	return nil
}

func enableProxy() error {
	/*
		echo "[Info]: creating proxy"
		${iptables} -t mangle -N PROXY

		for ignore in ${ignore_out_list[@]} ; do
		    ${iptables} -t mangle -A PROXY -o ${ignore} -j RETURN
		done

		if [ "${iptables}" = "ip6tables -w 100" ] ; then
		    for subnet6 in ${intranet6[@]} ; do
		        ${iptables} -t mangle -A PROXY -d ${subnet6} -j RETURN
		    done
		else
		    for subnet in ${intranet[@]} ; do
		        ${iptables} -t mangle -A PROXY -d ${subnet} -j RETURN
		    done
		fi

		# Bypass Xray itself
		${iptables} -t mangle -A PROXY -m owner --gid-owner ${inet_uid} -j RETURN
	*/
	return nil
}

func (t *linuxTProxy) DisableProxy() error {
	if runtime.GOOS == "android" {
		return nil
	}
	return nil
}

// No need for linux, but for Xray4Magisk
type linuxProxy struct{}

func makeIProxy() IProxy {
	return &linuxProxy{}
}

func (p *linuxProxy) ApplyProxy() {
	/*
		${iptables} -t mangle -A PROXY -p tcp -j MARK --set-mark ${mark_id}
		${iptables} -t mangle -A PROXY -p udp -j MARK --set-mark ${mark_id}
	*/
}
