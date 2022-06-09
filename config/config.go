package config

import (
	"github.com/go-ini/ini"
	"log"
	"os"
)

type ConfigList struct {
	ApiKey    string
	ApiSecret string
	LogFile   string
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Printf("faild to read file: %v", err)
		os.Exit(1)
	}

	Config = ConfigList{
		ApiKey:    cfg.Section("bitflyer").Key("api_key").String(),
		ApiSecret: cfg.Section("bitflyer").Key("api_secret").String(),
		LogFile:   cfg.Section("goflyer").Key("log_file").String(),
	}
}
