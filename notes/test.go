package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	coinPrice := map[string]float64{
		"btcusdt":  0.00,
		"ethusdt":  0.00,
		"dogeusdt": 0.00,
	}
	GetCoinPrice(coinPrice)
}

func GetCoinPrice(Price map[string]float64) {
	start := time.Now()
	ch := make(chan float64, 3)
	var wg sync.WaitGroup

	for k, _ := range Price {
		wg.Add()
		priceAPI := "https://api.huobi.pro/market/trade?symbol=" + k

		r, err := http.Get(priceAPI)
		if err != nil {
			log.Println("Get Coin Price http.Get ERROR ! ", err)
			continue
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
		Price[k] = coinprice.Tick.Data[0].Price
	}
	secs := time.Since(start).Seconds()
	log.Println(secs)
}
