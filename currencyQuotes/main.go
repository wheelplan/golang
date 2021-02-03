package main

import (
	"golang/currencyQuotes/price"
	"log"
	"time"

	"golang/currencyQuotes/mysql"
	//	"golang/currencyQuotes/sms"
	"golang/currencyQuotes/wechat"
)

func main() {

	expectedPrice := 1520.00

	for {
		coinPrice := map[string]float64{
			"btcusdt":  0.00,
			"ethusdt":  0.00,
			"dogeusdt": 0.00,
		}

		price.GetCoinPrice(coinPrice)

		for k, v := range coinPrice {

			if v > expectedPrice {
				wechat.Send(k, v)
				expectedPrice *= 1.01
			}

			log.Println(k, ": $", v)
			mysql.CoinMySQLData(k, v)
			time.Sleep(time.Duration(6) * time.Second)

		}

	}
}
