package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var (
	wsHost = "127.0.0.1:8080"
	wsPath = "/ws/conn"
)

type Client struct {
	Host string
	Path string
}

func NewWebsocketClient(host, path string) *Client {
	return &Client{
		Host: host,
		Path: path,
	}
}

// 给server发送输入的消息
func (this *Client) SendMessage() error {
	//
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	dialer := &websocket.Dialer{}
	conn, _, err := dialer.Dial("ws://"+this.Host+this.Path, http.Header{})
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer conn.Close() //关闭连接
	done := make(chan struct{})

	// 另外其一个goroutine处理接收消息
	go func() {
		defer close(done)
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Fatalf("read:", err)
			}
			log.Logf("recv msg from server values: %s", message)
		}
	}()
	inputFinish := make(chan bool, 1)
	inputFinish <- true
	log.Logf("input your message here:", err)
	var a string
	for {
		select {
		case <-interrupt:
			log.Fatalf("interrupt")
			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Fatalf("write close:", err)
				return err
			}
			return nil
		case <-inputFinish:
			// 处理输入
			fmt.Scanln(&a)
			if err = conn.WriteMessage(websocket.TextMessage, []byte(a)); err != nil {
				log.Fatal(err)
			}
			log.Logf("send msg client to server values:==>%s", a)
			inputFinish <- true
		default:
			fmt.Printf("no match case")
		}
	}
}

func msgHandler() {
	clientWrapper := NewWebsocketClient(wsHost, wsPath)
	if err := clientWrapper.SendMessage(); err != nil {
		log.Logf("SendMessage: errr%v", err)
	}
}

func main() {
	// 一个简历socket连接的客户端服务
	service := micro.NewService(
		micro.Name("go.micro.web.wsClient"),
	)
	service.Init()
	go msgHandler()
	if err := service.Run(); err != nil {
		log.Logf("Run: errr%v", err)
	}

}
