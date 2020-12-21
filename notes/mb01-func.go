package main

import (
	"fmt"
	"log"
	"time"
)

type serverOption struct {
	hostname string
	ip       string
	port     int
	timeout  time.Duration
	log      *log.Logger
}

func newOption() *serverOption {
	return &serverOption{
		hostname: "test-master",
		ip:       "192.168.1.20",
		port:     22,
		timeout:  time.Second * 10,
		log:      nil,
	}
}

func server(option *serverOption) { fmt.Println(option) }

func main() {
	opt := newOption()
	opt.timeout = time.Second * 5

	server(opt)

	opts := serverOption{
		hostname: "test-node1",
		ip:       "192.168.1.21",
		port:     22,
		timeout:  time.Second * 6,
		log:      nil,
	}
	opts.port = 3306
	server(&opts)
}
