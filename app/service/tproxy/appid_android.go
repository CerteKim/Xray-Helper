// +build android
package tproxy

import "log"

func init() {
	log.Println("Initialized appid module")
}

type androidProxy struct{}

func makeIProxy() IProxy {
	return &androidProxy{}
}

func (p *androidProxy) ApplyProxy() {
	switch model.ProxyMode() {
	case "ALL":
		/*
		 *	${iptables} -t mangle -A PROXY -p tcp -j MARK --set-mark ${mark_id}
		 *	${iptables} -t mangle -A PROXY -p udp -j MARK --set-mark ${mark_id}
		 */
	case "bypass":
		/*
		 *	for appid in ${appid_list[@]} ; do
		 *	   	${iptables} -t mangle -I PROXY -m owner --uid-owner ${appid} -j RETURN
		 *   done
		 *	${iptables} -t mangle -A PROXY -p tcp -j MARK --set-mark ${mark_id}
		 *	${iptables} -t mangle -A PROXY -p udp -j MARK --set-mark ${mark_id}
		 */
	case "":
		/*
		 *	for appid in ${appid_list[@]} ; do
		 *    	${iptables} -t mangle -A PROXY -p tcp -m owner --uid-owner ${appid} -j MARK --set-mark ${mark_id}
		 *    	${iptables} -t mangle -A PROXY -p udp -m owner --uid-owner ${appid} -j MARK --set-mark ${mark_id}
		 *   done
		 */
	}
}
