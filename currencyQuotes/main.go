package main

import (
	"golang/currencyQuotes/mysql"
	"golang/currencyQuotes/price"
	"log"
	"time"

	//	"golang/currencyQuotes/sms"
	"golang/currencyQuotes/wechat"
)

func main() {

	coinPrice := map[string]float64{
		"btcusdt":  0.00,
		"ethusdt":  0.00,
		"dogeusdt": 0.00,
	}
	expectedPrice := map[string]float64{
		"btcusdt":  37000.00,
		"ethusdt":  1530.00,
		"dogeusdt": 0.034,
	}

	for {
		price.GetCoinPrice(coinPrice)

		for k, v := range coinPrice {

			if v > expectedPrice[k] {
				wechat.Send(k, v)
				expectedPrice[k] *= 1.01
			}

			log.Println(k, ": $", v)
			mysql.CoinMySQLData(k, v)
		}

		time.Sleep(time.Duration(6) * time.Second)
	}
}
