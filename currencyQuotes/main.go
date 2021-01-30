package main                                                                                                                                     

import (
    "log"
    "strconv"
    "time"
    "golang/currencyQuotes/mysql"
    "golang/currencyQuotes/messages"
)

func main() {

    corpid := "wwd327d0ea50c0dae0"
    appsecret := "1kJw7t38aUxb3doG7IaLGVwOsDEuDHOMFUO2Z0xa_lI"
    agentid := 1000008
    toparty := "1"

    expectedPrice := 1314.00

    for {
        accessToken := messages.GetToken(corpid, appsecret)
        coinPrice := messages.GetCoinPrice("ethusdt")
        msg := "ETH Coin Price is $" + strconv.FormatFloat(coinPrice, 'f', 2, 64)
        message := []string{msg}

        if coinPrice < expectedPrice {
            messages.SendMSG(accessToken, agentid, toparty, message)
            expectedPrice = expectedPrice * 0.99
        }

        log.Println("eth: $", coinPrice)
        mysql.CoinMySQLData(coinPrice)
        time.Sleep(time.Duration(6) * time.Second)
    }
}

