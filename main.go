package main

import (
	"log"
	"goflyer/bitflyer"
	"goflyer/config"
	"goflyer/utils"

	"fmt"
)

func init() {
	utils.LoggingSettings(config.Config.LogFile)
}

func main() {
	log.Println("test")
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)
	fmt.Println(apiClient.GetBalance())
}
