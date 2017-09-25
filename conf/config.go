package conf

import (
	"log"
	"encoding/json"
	"os"
)

type Configuration struct {
	Address      string
	ReadTimeOut  int64
	WriteTimeOut int64
	Static       string
	DBusername   string
	DBpassword   string
	DBname       string
}

var Config Configuration

func init() {
	loadConfig()
}

func loadConfig() {
	file, err := os.Open("./conf/config.json")
	if err != nil {
		log.Fatalln("Cannot open Config file", err)
	}
	decoder := json.NewDecoder(file)
	Config = Configuration{}
	err = decoder.Decode(&Config)
	if err != nil {
		log.Fatalln("Cannot get configuration from file", err)
	}
}
