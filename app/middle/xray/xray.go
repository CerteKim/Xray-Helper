package xray

import (
	"encoding/json"
	"log"

	"github.com/certekim/xray-helper/app/middle"
)

type ConfigObj map[string]interface{}

//type ConfigObjs []ConfigObj

/*******************************************************
  path:
	Xray4Magisk default: /data/xray/helper/jsons/confs
	Linux default: /usr/local/etc/xray
  tag include:
	"base"
	"dns"
	"proxy"
	"routing"
	"xray-api"
*******************************************************/

func Write(tag string, in map[string]interface{}) {
	if err := middle.JsonDB.Write("confs", tag, in); err != nil {
		log.Println("Error", err)
	}
}

func Read(tag string) []byte {
	var ConfigObj ConfigObj
	if err := middle.JsonDB.Read("confs", tag, &ConfigObj); err != nil {
		log.Println("Error", err)
	}
	if ret, err := json.Marshal(ConfigObj); err != nil {
		log.Println("Error", err)
	} else {
		return ret
	}
	return nil
}

func Del(tag string) {
	if err := middle.JsonDB.Delete("confs", tag); err != nil {
		log.Println("Error", err)
	}
}
