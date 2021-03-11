package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"sort"
	"time"
)

const (
	SecretKey = "8a03c923-805388e0-b57aa7b0-af1ce"
)

func main() {
<<<<<<< HEAD

	params := map[string]string {
		"SignatureMethod":	"HmacSHA256",
		"SignatureVersion":	"2",
		"Timestamp":     	time.Now().UTC().Format("2006-01-02T15:04:05"),
		"AccessKeyId":   	"ffde1a6e-31848ea0-45e77dbf-mn8ikls4qg",
		"accountType":		"spot",
		"valuationCurrency": 	"CNY",
=======
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
>>>>>>> 45585450458f8c77f15719cd10c89264b0db6a35
	}


	var buf bytes.Buffer
	a := "GET\n" + "api.huobi.pro\n" + "/v2/account/asset-valuation\n"

	keys := make([]string, 0, len(params))
	for k, _ := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for i := range keys {
		k := keys[i]
		buf.WriteString(k)
		buf.WriteString("=")
		buf.WriteString(url.QueryEscape(params[k]))
		buf.WriteString("&")
	}
	buf.Truncate(buf.Len() - 1)

	hashed := hmac.New(sha256.New, []byte(SecretKey))
	var data []byte = []byte(a)
	hashed.Write(append(data, buf.Bytes()...))

	r, err := http.Get("https://api.huobi.pro/v2/account/asset-valuation?" +
		string(buf.Bytes()) +
		"&Signature=" +
		url.QueryEscape(base64.StdEncoding.EncodeToString(hashed.Sum(nil))))

	if err != nil {
		fmt.Println(err)
		return
	}

	io.Copy(os.Stdout, r.Body)
	if err := r.Body.Close(); err != nil {
		fmt.Println(err)
	}
}
