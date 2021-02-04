package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	start := time.Now()
	Price := map[string]float64{
		"btcusdt":  0.00,
		"ethusdt":  0.00,
		"dogeusdt": 0.00,
	}
	ch := make(chan string)
	for currency, _ := range Price {
		priceAPI := "https://api.huobi.pro/market/trade?symbol=" + currency
		go fetch(priceAPI, ch)
	}
	for range Price {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) float64 {
	start := time.Now()

	r, err := http.Get(url)
	if err != nil {
		log.Println("Get Coin Price http.Get ERROR ! ", err)
	}

	defer r.Body.Close()
	type coinPrice struct {
		Ch     string `json:"ch"`
		Status string `json:"status"`
		//	Ts     time.Time `json:"ts"`
		Tick struct {
			ID int `json:"id"`
			//	Ts1  time.Time `json:"ts"`
			Data []struct {
				//Id int `json:"id"`
				//Ts2       time.Time `json:"ts"`
				//TradeId int `json:"trade-id"`
				//Amount    float64 `json:"amount"`
				Price     float64 `json:"price"`
				Direction string  `json:"direction"`
			} `json:"data"`
		} `json:"tick"`
	}

	var coinprice coinPrice
	jsonERR := json.NewDecoder(r.Body).Decode(&coinprice)
	if jsonERR != nil {
		log.Fatalln("Get Coin Price json.NewDecoder.Decode ERR ! ", jsonERR)
	}

	// return coinprice.Tick.Data[0].Price
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, coinprice.Tick.Data[0].Price, url)

	return coinprice.Tick.Data[0].Price
}
