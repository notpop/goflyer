package main

import (
	"goflyer/bitflyer"
	"goflyer/config"
	"goflyer/utils"
	"log"
	"time"

	"fmt"
)

func init() {
	utils.LoggingSettings(config.Config.LogFile)
}

func main() {
	log.Println("start")
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)

	tickerChannel := make(chan bitflyer.Ticker)
	go apiClient.GetRealTimeTicker(config.Config.ProductCode, tickerChannel)
	for ticker := range tickerChannel {
		fmt.Println(ticker)
		fmt.Println(ticker.GetMidPrice())
		// fmt.Println(ticker.DateTime())
		fmt.Println(ticker.TruncateDateTime(time.Second))
		// fmt.Println(ticker.TruncateDateTime(time.Minute))
		// fmt.Println(ticker.TruncateDateTime(time.Hour))
	}

	log.Println("end")
}
