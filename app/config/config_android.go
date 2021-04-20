// +build android !linux

package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func ReadConf() {
	var data []byte
	var err error
	if path := os.Getenv("XRAY_HELPER_CONFIG_PATH"); path != "" {
		data, err = ioutil.ReadFile(path + "/config.json")
		if err != nil {
			log.Printf("read file error on %s/config.json : ", path)
			log.Println(err)
		} else {
			log.Printf("read file on %s/config.json : ", path)
		}
	} else {
		path, _ := os.Getwd()
		data, err = ioutil.ReadFile(path + "/config.json")
		if err != nil {
			log.Printf("read file error on %s/config.json : ", path)
			log.Println(err)
		} else {
			log.Printf("read file on %s/config.json : ", path)
		}
	}
	err = json.Unmarshal(data, &Conf)
	if err != nil {
		log.Println(err)
	}
}
