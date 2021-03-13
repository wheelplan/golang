package main

import (
	"encoding/json"
	"fmt"
	//"golang/currencyQuotes/msg"
	//"golang/currencyQuotes/mysql"
	//"golang/currencyQuotes/price"
	//"golang/currencyQuotes/wechat"
	"log"
	"net/http"
	//"os"
	//"strconv"
)
	var coinPrice = map[string]float64{
		"btcusdt":  0.00,
		"ethusdt":  0.00,
		"htusdt": 0.00,
		"dogeusdt": 0.00,
	}

type coin string

func main() {
	// 设置路由，如果访问/，则调用index方法
	var eth coin = "ethusdt"
	var doge coin = "dogeusdt"
	var btc coin = "btcusdt"
	var ht coin = "htusdt"
	http.HandleFunc("/eth", eth.price)
	http.HandleFunc("/ht", ht.price)
	http.HandleFunc("/btc", btc.price)
	http.HandleFunc("/doge", doge.price)

	// 启动web服务，监听9090端口
	//err := http.ListenAndServe(":9090", nil)
	//if err != nil {
	//	log.Fatal("ListenAndServe: ", err)
	//}
	log.Fatal(http.ListenAndServe(":9090", nil))

}

// w表示response对象，返回给客户端的内容都在对象里处理
// r表示客户端请求对象，包含了请求头，请求参数等等
func (c coin) price(w http.ResponseWriter, r *http.Request) {
	// 往w里写入内容，就会在浏览器里输出
	GetCoinPrice(coinPrice)
	fmt.Fprintf(w, "%f", coinPrice[string(c)])
	//fmt.Fprintf(w, "%q", strconv.FormatFloat(coinPrice[string(c)], 'f', 4, 64))
}


func GetCoinPrice(Price map[string]float64) {
	for k, _ := range Price {
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
}
