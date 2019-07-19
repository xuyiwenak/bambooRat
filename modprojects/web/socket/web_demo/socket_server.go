package main

import (
	"github.com/gorilla/websocket"
	"github.com/micro/go-micro/web"
	"log"
	"net/http"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func main() {
	// New web service
	service := web.NewService(
		web.Name("go.micro.web.websocket"),
	)

	if err := service.Init(); err != nil {
		log.Fatal("Init", err)
	}

	// 这是返回一个静态网页，如果直接建立连接可以忽略
	service.Handle("/websocket/", http.StripPrefix("/websocket/", http.FileServer(http.Dir("html"))))

	// 直接建立一个websocket连接
	service.HandleFunc("/websocket/hi", hi)

	if err := service.Run(); err != nil {
		log.Fatal("Run: ", err)
	}
}

func hi(w http.ResponseWriter, r *http.Request) {

	c, err := upGrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("upgrade: %s", err)
		return
	}

	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}

		log.Printf("recv: %s", message)

		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
