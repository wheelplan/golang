package main

import (
	"log"
	"os"
	"strconv"
	"time"
)

func main() {

	sign := os.Args[1]
	if len(os.Args) != 5 {
		log.Fatalln("Please input: ./msg '> or <' 'btcPrice' 'ethPrice' 'dogePrice' .")
	}
	expectedBTC, _ := strconv.ParseFloat(os.Args[2], 64)
	expectedETH, _ := strconv.ParseFloat(os.Args[3], 64)
	expectedDOGE, _ := strconv.ParseFloat(os.Args[4], 64)

	coinPrice := map[string]float64{
		"btcusdt":  0.00,
		"ethusdt":  0.00,
		"dogeusdt": 0.00,
	}
	expectedPrice := map[string]float64{
		"btcusdt":  expectedBTC,
		"ethusdt":  expectedETH,
		"dogeusdt": expectedDOGE,
	}

	for {
		GetCoinPrice(coinPrice)

		for k, v := range coinPrice {

			if sign == ">" {
				if v > expectedPrice[k] {
					wechatSend(k, v)
					expectedPrice[k] *= 1.01
				}

				log.Println(k, ": $", v)

			} else if sign == "<" {
				if v < expectedPrice[k] {
					wechatSend(k, v)
					expectedPrice[k] *= 0.99
				}

				log.Println(k, ": $", v)
			}
		}

		CoinMySQLData(coinPrice)
		time.Sleep(time.Duration(6) * time.Second)
	}
}
