package main

import (
	"golang/currencyQuotes/price"
	"log"
	"time"

	"golang/currencyQuotes/mysql"
	"golang/currencyQuotes/sms"
	"golang/currencyQuotes/wechat"
	"strconv"
)

func main() {

	expectedPrice := 1440.00

	for {
		coinPrice := price.GetCoinPrice("ethusdt")

		if coinPrice > expectedPrice {
			wechat.Send(strconv.FormatFloat(coinPrice, 'f', 2, 64))
			sms.Send(strconv.FormatFloat(coinPrice, 'f', 2, 64))
			expectedPrice *= 1.01
		}

		log.Println("eth: $", coinPrice)
		mysql.CoinMySQLData(coinPrice)
		time.Sleep(time.Duration(6) * time.Second)
	}
}
