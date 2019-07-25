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
		web.Name("go.micro.web.ws"),
	)

	if err := service.Init(); err != nil {
		log.Fatal("Init", err)
	}
	// websocket 连接接口 web.name注册根据.分割路由路径，所以注册的路径要和name对应上
	service.HandleFunc("/ws/conn", conn)

	if err := service.Run(); err != nil {
		log.Fatal("Run: ", err)
	}
}
func conn(w http.ResponseWriter, r *http.Request) {

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
		resStr := string(message)
		resStr += "<===server"
		err = c.WriteMessage(mt, []byte(resStr))
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
