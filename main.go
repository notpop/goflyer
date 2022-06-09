package main

import (
	"goflyer/app/controllers"
	// "goflyer/app/models"
	// "goflyer/bitflyer"
	"goflyer/config"
	"goflyer/utils"
	"log"
	// "time"
	// "fmt"
)

func init() {
	utils.LoggingSettings(config.Config.LogFile)
}

func main() {
	log.Println("start")
	// apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)

	// tickerChannel := make(chan bitflyer.Ticker)
	// go apiClient.GetRealTimeTicker(config.Config.ProductCode, tickerChannel)
	// for ticker := range tickerChannel {
	// 	fmt.Println(ticker)
	// 	fmt.Println(ticker.GetMidPrice())
	// 	fmt.Println(ticker.DateTime())
	// 	fmt.Println(ticker.TruncateDateTime(time.Second))
	// 	fmt.Println(ticker.TruncateDateTime(time.Minute))
	// 	fmt.Println(ticker.TruncateDateTime(time.Hour))

	// 	// 試しなので一件だけ取得したら終了
	// 	break
	// }

	// 実際に売りを入れてそのオーダーを監視する
	// order := &bitflyer.Order {
	// 	ProductCode: config.Config.ProductCode,
	// 	ChildOrderType: "MARKET",
	// 	Side: "SELL",
	// 	Size: 0.001,
	// 	MinuteToExpires: 1,
	// 	TimeInForce: "GTC",
	// }
	// res, _ := apiClient.SendOrder(order)
	// orderId := res.ChildOrderAcceptanceID
	// params := map[string]string{
	// 	"product_code": config.Config.ProductCode,
	// 	"child_order_acceptance_id": orderId,
	// }
	// r, _ := apiClient.ListOrder(params)
	// fmt.Println(r)

	// fmt.Println(models.DbConnection)
	controllers.StreamInjestionData()
	controllers.StartWebServer()

	log.Println("end")
}
