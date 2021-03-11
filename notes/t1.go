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

	params := map[string]string {
		"SignatureMethod":	"HmacSHA256",
		"SignatureVersion":	"2",
		"Timestamp":     	time.Now().UTC().Format("2006-01-02T15:04:05"),
		"AccessKeyId":   	"ffde1a6e-31848ea0-45e77dbf-mn8ikls4qg",
		"accountType":		"spot",
		"valuationCurrency": 	"CNY",
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
