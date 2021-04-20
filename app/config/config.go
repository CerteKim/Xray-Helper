package config

import (
	"log"
)

var Conf map[string]interface{}

func init() {
	ReadConf()
	log.Println("Listen port :", Conf["port"])
	log.Println("Connect port :", Conf["xray"])
	log.Println("Read config finished !")
}
