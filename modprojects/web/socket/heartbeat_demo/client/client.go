package main

import (
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	"golang.org/x/net/websocket"
	"net/url"
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

func (this *Client) SendMessage(body []byte) error {
	u := url.URL{Scheme: "ws", Host: this.Host, Path: this.Path}

	ws, err := websocket.Dial(u.String(), "", "http://"+this.Host+"/")
	defer ws.Close() //关闭连接
	if err != nil {
		log.Fatal(err)
		return err
	}
	_, err = ws.Write(body)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
func main() {
	// 一个简历socket连接的客户端服务
	service := micro.NewService(
		micro.Name("go.micro.web.wsClient"),
	)
	service.Init()
	clientWrapper := NewWebsocketClient(wsHost, wsPath)
	greeterStr := "hello"

	if err := clientWrapper.SendMessage([]byte(greeterStr)); err != nil {
		fmt.Printf("SendMessage: errr%v", err)
	}
	if err := service.Run(); err != nil {
		fmt.Printf("Run: errr%v", err)
	}
}
