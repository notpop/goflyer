package controllers

import (
	"goflyer/app/models"
	"goflyer/bitflyer"
	"goflyer/config"
	"log"
)

func StreamInjestionData() {
	var tickerChannel = make(chan bitflyer.Ticker)
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)
	go apiClient.GetRealTimeTicker(config.Config.ProductCode, tickerChannel)
	go func() {
		for ticker := range tickerChannel {
			log.Printf("action=StreamInjestionData, %v", ticker)
			for _, duration := range config.Config.Durations {
				isCreated := models.CreateCandleWithDuration(ticker, ticker.ProductCode, duration)
				if isCreated && duration == config.Config.TradeDuration {
					log.Printf("%v", "TODO")
				}
			}
		}
	}()
}
