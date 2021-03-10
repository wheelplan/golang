package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	SecretKey = "8a03c923-805388e0-b57aa7b0-af1ce"
)

func main() {
	AccessKey := "ffde1a6e-31848ea0-45e77dbf-mn8ikls4qg"
	timestamp := time.Now().Format("2006-01-02T15:04:05")

	msg := "AccessKeyId=" + AccessKey +
		"&SignatureMethod=HmacSHA256" +
		"&SignatureVersion=2" +
		"&Timestamp=" + timestamp

	message := "GET\n" + "api.huobi.pro\n" + "/v2/account/asset-valuation\n" + msg

	r, err := http.Get("https://api.huobi.pro/v2/account/asset-valuation?" + msg +
		"&Signature=" + ComputeHmacSha256(message, SecretKey))
	if err != nil {
		log.Println("Get Coin Price http.Get ERROR ! ", err)
	}
	defer r.Body.Close()
	fmt.Println(r.Body)
}

func ComputeHmacSha256(message, secret string) string {

	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return base64.URLEncoding.EncodeToString(h.Sum(nil))
}
