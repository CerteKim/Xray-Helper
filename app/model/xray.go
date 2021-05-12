package model

import (
	"encoding/json"
	"log"
)

type ConfigObj map[string]interface{}

//type ConfigObjs []ConfigObj

/*******************************************************
  path:
	Xray4Magisk default: /data/adb/xray/xrayd/confs
	Linux default: /usr/local/etc/xray
  tag include:
	"base"
	"dns"
	"proxy"
	"routing"
	"xray-api"
	"xray-web"
*******************************************************/

func XrayWrite(tag string, in map[string]interface{}) {
	if err := JsonDB.Write("confs", tag, in); err != nil {
		log.Println("Error", err)
	}
}

func XrayRead(tag string) []byte {
	var ConfigObj ConfigObj
	if err := JsonDB.Read("confs", tag, &ConfigObj); err != nil {
		log.Println("Error", err)
	}
	if ret, err := json.Marshal(ConfigObj); err != nil {
		log.Println("Error", err)
	} else {
		return ret
	}
	return nil
}

func XrayDel(tag string) {
	if err := JsonDB.Delete("confs", tag); err != nil {
		log.Println("Error", err)
	}
}
