package main

import (
	"github.com/micro/go-micro/transport"
	"github.com/micro/go-micro/util/log"
)

//transport模块的接口主要是用于服务之间连接。它使用基于连接的socket send/recv语义和
//有各种实现；http, grpc, memory, quic. 默认使用http.

func main() {
	tr := transport.NewTransport()

	l, err := tr.Listen("127.0.0.1:54321")
	log.Logf("listening on port:%s ...", l.Addr())
	if err != nil {
		log.Fatalf("Unexpected listen err: %v", err)
	}
	defer l.Close()

	fn := func(sock transport.Socket) {
		defer sock.Close()
		for {
			var m transport.Message
			if err := sock.Recv(&m); err != nil {
				log.Fatalf("rev", err)
				return
			}

			if err := sock.Send(&m); err != nil {
				log.Fatalf("send", err)
				return
			}
		}
	}

	done := make(chan bool)

	go func() {
		if err := l.Accept(fn); err != nil {
			select {
			case <-done:
			default:
				log.Fatalf("Unexpected accept err: %v", err)
			}
		}
	}()

	c, err := tr.Dial(l.Addr())
	if err != nil {
		log.Fatalf("Unexpected dial err: %v", err)
	}
	defer c.Close()

	m := transport.Message{
		Header: map[string]string{
			"Content-Type": "application/json",
		},
		Body: []byte(`{"message": "Hello World"}`),
	}

	if err := c.Send(&m); err != nil {
		log.Fatalf("Unexpected send err: %v", err)
	}

	var rm transport.Message

	if err := c.Recv(&rm); err != nil {
		log.Fatalf("Unexpected recv err: %v", err)
	}

	if string(rm.Body) != string(m.Body) {
		log.Fatalf("Expected %v, got %v", m.Body, rm.Body)
	}
	close(done)
}
