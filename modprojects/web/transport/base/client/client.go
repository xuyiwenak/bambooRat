package main

import (
	"github.com/micro/go-micro/transport"
	"github.com/micro/go-micro/util/log"
	"time"
)

//transport模块的接口主要是用于服务之间连接。它使用基于连接的socket send/recv语义和
//有各种实现；http, grpc, memory, quic. 默认使用http.

var (
	addr = "127.0.0.1:54321"
)

func msghandler(i int) transport.Message {
	// send message
	return transport.Message{
		Header: map[string]string{
			"Content-Type": "application/json",
		},
		Body: []byte("server send msg seq:" + string(i)),
	}
}
func main() {

	tr := transport.NewTransport()

	done := make(chan bool)

	// client connection
	client, err := tr.Dial(addr)
	if err != nil {
		log.Fatalf("Unexpected dial err: %v", err)
	}
	counter := 0
	msg := msghandler(counter)
	ticker := time.NewTicker(time.Second * 5)
	defer ticker.Stop()
	// send message
	for range ticker.C {
		counter++
		log.Logf("count:%d", counter)
		if err := client.Send(&msg); err != nil {
			log.Fatalf("Unexpected send err: %v", err)
		}
	}

	/*
		go func() {
			var rm transport.Message
			// receive message
			if err := client.Recv(&rm); err != nil {
				log.Fatalf("Unexpected recv err: %v", err)
			}
		}()
	*/
	client.Close()
	log.Logf("client ready to close!")
	// finish
	close(done)
}
