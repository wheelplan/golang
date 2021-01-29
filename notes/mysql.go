package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"goNotes/wechatMessageAPI"
	"time"
)

const (
	UserName = "lucky"
	PassWord = "linux2020"
	NetWork  = "tcp"
	Server   = "yuiya.top"
	Port     = 3306
	Database = "eth"
)

type Price struct {
	Time  time.Time `json:"time"`
	Price float64   `json:"price"`
}

func main() {

	conn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", UserName, PassWord, NetWork, Server, Port, Database)
	DB, err := sql.Open("mysql", conn)
	if err != nil {
		fmt.Println("connection to mysql failed:", err)
		return
	}

	defer DB.Close()
	DB.SetConnMaxLifetime(100 * time.Second)
	DB.SetMaxOpenConns(100)
	InsertData(DB)

}

func InsertData(DB *sql.DB) {
	ethPrice := wechatMessageAPI.GetCoinPrice("ethusdt")
	result, err := DB.Exec("insert INTO price(time,price) values(?,?)", time.Now(), ethPrice)
	if err != nil {
		fmt.Printf("Insert data failed,err:%v", err)
		return
	}
	lastInsertID, err := result.LastInsertId() //获取插入数据的自增ID
	if err != nil {
		fmt.Printf("Get insert id failed,err:%v", err)
		return
	}
	fmt.Println("Insert data id:", lastInsertID)

	rowsaffected, err := result.RowsAffected() //通过RowsAffected获取受影响的行数
	if err != nil {
		fmt.Printf("Get RowsAffected failed,err:%v", err)
		return
	}
	fmt.Println("Affected rows:", rowsaffected)
}
