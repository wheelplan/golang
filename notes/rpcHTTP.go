package main

import (
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	rpc.RegisterName("He", new(HelloService))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			ReadCloser: r.Body,
			Writer:     w,
		}

		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})

	http.ListenAndServe(":10087", nil)
}

type HelloService struct{}

func (p *HelloService) Hi(request string, reply *string) error {
	*reply = "hi," + request
	return nil
}
