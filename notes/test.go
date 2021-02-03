package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	//rpc.RegisterName("He", new(Hee))
	//
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	var conn io.ReadWriteCloser = struct {
	//		io.Writer
	//		io.ReadCloser
	//	}{
	//		ReadCloser: r.Body,
	//		Writer:     w,
	//	}
	//
	//	rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	//})
	//
	//http.ListenAndServe(":10087", nil)
	price := map[string]float64{
		"btcusdt":  0.00,
		"ethusdt":  0.00,
		"dogeusdt": 0.00,
	}
	fmt.Print(reflect.TypeOf(price), price["btcusdt"])
}

func add(s string) {
	adb := s + "da"
	fmt.Println(adb)

	f := 3.1415
	fmt.Println(reflect.TypeOf(strconv.FormatFloat(f, 'f', 2, 64)))
}

type Hee struct{}

func (p *Hee) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}
