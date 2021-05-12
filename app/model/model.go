package model

import (
	"log"

	scribble "github.com/nanobox-io/golang-scribble"
	"github.com/spf13/viper"
)

var JsonDB *scribble.Driver

func InitDB() {
	var err error
	path := viper.GetString("xray.confdir")
	JsonDB, err = scribble.New(path, nil)
	if err != nil {
		log.Println("storage pool:")
		log.Println("	Error: ", err)
	} else {
		log.Println("storage pool connected: " + path)
	}
}
