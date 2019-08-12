package main

import (
	"github.com/micro/go-micro/transport"
	"github.com/micro/go-micro/util/log"
	"sync"
)

//transport模块的接口主要是用于服务之间连接。它使用基于连接的socket send/recv语义和
//有各种实现；http, grpc, memory, quic. 默认使用http.

func main() {

	tr := transport.NewTransport()

	// server listen
	l, err := tr.Listen("127.0.0.1:54321")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	// socket func
	fn := func(sock transport.Socket) {
		defer sock.Close()

		for {
			var m transport.Message
			if err := sock.Recv(&m); err != nil {
				log.Fatal(err)
				return
			}

			if err := sock.Send(&m); err != nil {
				log.Fatal(err)
				return
			}
		}
	}

	done := make(chan bool)

	// accept connections
	go func() {
		if err := l.Accept(fn); err != nil {
			select {
			case <-done:
			default:
				log.Fatalf("Unexpected accept err: %v", err)
			}
		}
	}()

	m := transport.Message{
		Header: map[string]string{
			"Content-Type": "application/json",
		},
		Body: []byte(`{"message": "Hello World"}`),
	}

	// client connection
	client, err := tr.Dial(l.Addr())
	if err != nil {
		log.Fatalf("Unexpected dial err: %v", err)
	}

	send := func(c transport.Client) {
		// send message
		if err := c.Send(&m); err != nil {
			log.Fatalf("Unexpected send err: %v", err)
		}

		var rm transport.Message
		// receive message
		if err := c.Recv(&rm); err != nil {
			log.Fatalf("Unexpected recv err: %v", err)
		}
	}

	// warm
	for i := 0; i < 10; i++ {
		send(client)
	}

	client.Close()

	ch := make(chan int, 2)

	var wg sync.WaitGroup
	wg.Add(2)

	for i := 0; i < 2; i++ {
		go func() {
			cl, err := tr.Dial(l.Addr())
			if err != nil {
				log.Fatalf("Unexpected dial err: %v", err)
			}
			defer cl.Close()

			for range ch {
				send(cl)
			}

			wg.Done()
		}()
	}

	for i := 0; i < 2; i++ {
		ch <- i
	}
	close(ch)

	wg.Wait()
	log.Logf("client ready to close!")
	// finish
	close(done)

}
