package price

import (
	"encoding/json"
	"log"
	"net/http"
)

func GetCoinPrice(coin string) float64 {
	priceAPI := "https://api.huobi.pro/market/trade?symbol=" + coin

	r, err := http.Get(priceAPI)
	if err != nil {
		log.Fatalln("Get Coin Price http.Get ERROR ! ", err)
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

	return coinprice.Tick.Data[0].Price
}
