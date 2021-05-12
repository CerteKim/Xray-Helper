package config

import (
	"log"
	"os"
	"runtime"

	"github.com/spf13/viper"
)

func LoadConfig() {
	log.Println("Current platform: ", runtime.GOOS)
	viper.SetConfigName("xrayd")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(os.Getenv("XRAYD_CONFIG_DIR"))
	viper.AddConfigPath("/etc/xrayd/")
	viper.AddConfigPath("/data/adb/xray/")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Fatal error config file: ", err)
	}
	log.Println("Using config file: ", viper.GetViper().ConfigFileUsed())
}
