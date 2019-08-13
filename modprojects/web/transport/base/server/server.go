package main

import (
	"github.com/micro/go-micro/transport"
	"github.com/micro/go-micro/util/log"
)

//transport模块的接口主要是用于服务之间连接。它使用基于连接的socket send/recv语义和
//有各种实现；http, grpc, memory, quic. 默认使用http.
var (
	addr = "127.0.0.1:54321"
)

func main() {
	tr := transport.NewTransport()

	l, err := tr.Listen(addr)
	log.Logf("listening on port:%s ...", addr)
	if err != nil {
		log.Fatalf("Unexpected listen err: %v", err)
	}
	defer l.Close()

	done := make(chan bool)
	// socket func
	fn := func(sock transport.Socket) {
		defer sock.Close()

		for {
			var msg transport.Message
			if err := sock.Recv(&msg); err != nil {
				return
			}
			log.Logf("recv msg:%s", string(msg.Body))
			if err := sock.Send(&msg); err != nil {
				return
			}
		}
	}
	// 这里先阻塞，如果不想的话启动一个协程单独跑
	if err := l.Accept(fn); err != nil {
		select {
		case <-done:
		default:
			log.Fatalf("Unexpected accept err: %v", err)
		}
	}
	close(done)
}
