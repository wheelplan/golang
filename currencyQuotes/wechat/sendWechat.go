package wechat

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func Send(coinPrice string) {
	corpid := "wwd327d0ea50c0dae0"
	appsecret := "sm8wQYy2F8R8QTy67hcm8-rA-nbp29kjzVcOSwTdbjM"
	agentid := 1000007
	toparty := "1"

	accessToken := GetToken(corpid, appsecret)

	message1 := "ETH Coin Price is $" + coinPrice
	//message2 := ""
	message := []string{message1}

	sendURL := "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=" + accessToken

	var messages string
	for _, arg := range message[:] {
		messages += arg + "\n"
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
		}{Content: messages},
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
