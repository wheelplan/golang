package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {

	/*	corpid := "wwd327d0ea50c0dae0"
		appsecret := "1kJw7t38aUxb3doG7IaLGVwOsDEuDHOMFUO2Z0xa_lI"
		agentid := 1000008
		toparty := os.Args[1]

		accessToken := GetToken(corpid, appsecret)
		SendMSG(accessToken, agentid, toparty)*/

	fmt.Println(GetCoinPrice("ethusdt"))

}

func GetToken(corpid, appsecret string) string {

	tokenURL := "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=" + corpid + "&corpsecret=" + appsecret

	r, err := http.Get(tokenURL)
	if err != nil {
		log.Fatalln("Get Token http.Get ERROR ! ", err)
	}

	defer r.Body.Close()

	type tokenBody struct {
		ErrCode     int    `json:"errcode"`
		ErrMsg      string `json:"errmsg"`
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	}

	var token tokenBody
	jsonERR := json.NewDecoder(r.Body).Decode(&token)
	if jsonERR != nil {
		log.Fatalln("Get Token json.NewDecoder.Decode ERR ! ", jsonERR)
	}

	return token.AccessToken
}

func SendMSG(accessToken string, agentid int, toparty string) {

	sendURL := "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=" + accessToken

	var message string
	for _, arg := range os.Args[2:] {
		message += arg + "\n"
	}

	type Params struct {
		Toparty string `json:"toparty"`
		Msgtype string `json:"msgtype"`
		Agentid int    `json:"agentid"`
		Text    struct {
			Content string `json:"content"`
		} `json:"text"`
		Safe int `json:"safe"`
	}

	params := &Params{
		Toparty: toparty,
		Msgtype: "text",
		Agentid: agentid,
		Text: struct {
			Content string `json:"content"`
		}{Content: message},
		Safe: 0,
	}

	paramsJSON, _ := json.Marshal(params)

	req, err := http.Post(sendURL, "application/json", bytes.NewBuffer(paramsJSON))
	if err != nil {
		log.Fatalln("SendMSG HTTP POST ERR ! ", err)
	}

	defer req.Body.Close()

	errCode := req.Header["Error-Code"][0]
	if errCode != "0" {
		log.Println("status:", req.Status)
		log.Println("response Header:", req.Header)
		body, _ := ioutil.ReadAll(req.Body)
		log.Println("response Body:", string(body))
	}
}

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
			ID string `json:"id"`
			//	Ts1  time.Time `json:"ts"`
			Data struct {
				Id string `json:"id"`
				//Ts2       time.Time `json:"ts"`
				TradeId   string  `json:"trade-id"`
				Amount    float64 `json:"amount"`
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

	return coinprice.Tick.Data.Price
}
