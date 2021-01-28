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
	corpid := "wwd327d0ea50c0dae0"
	appsecret := "1kJw7t38aUxb3doG7IaLGVwOsDEuDHOMFUO2Z0xa_lI"
	agentid := 1000008
	toparty := os.Args[1]
	accessToken := GetToken(corpid, appsecret)
	SendMSG(accessToken, agentid, toparty)
}

func GetToken(corpid, appsecret string) string {
	tokenURL := "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=" + corpid + "&corpsecret=" + appsecret

	r, err := http.Get(tokenURL)
	if err != nil {
		fmt.Println("http.Get bad ! ", err)
		return "http.GET ERR"
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("r.Body ERR ! ", err)
	}

	type tokenData struct {
		ErrCode     int    `json:"errcode"`
		ErrMsg      string `json:"errmsg"`
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	}

	var token tokenData
	jsonERR := json.Unmarshal(body, &token)
	if err != nil {
		fmt.Println("json.Unmarshal ERR ! ", jsonERR)
	}

	return token.AccessToken
}

func SendMSG(accessToken string, agentid int, toparty string) {
	sendMSG := "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=" + accessToken
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

	params := Params{
		Toparty: toparty,
		Msgtype: "text",
		Agentid: agentid,
		Text: struct {
			Content string `json:"content"`
		}{Content: message},
		Safe: 0,
	}

	paramsJSON, _ := json.Marshal(params)

	req, err := http.NewRequest("POST", sendMSG, bytes.NewBuffer(paramsJSON))
	if err != nil {
		fmt.Println("HTTP POST ERR ! ", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("charset", "UTF-8")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	errCode := resp.Header["Error-Code"][0]
	if errCode != "0" {
		log.Println("status:", resp.Status)
		log.Println("response Header:", resp.Header)
		body, _ := ioutil.ReadAll(resp.Body)
		log.Println("response Body:", string(body))
	}
}
