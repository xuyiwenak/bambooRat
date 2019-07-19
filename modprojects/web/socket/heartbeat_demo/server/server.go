package main

import (
	"bytes"
	"encoding/binary"
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
		web.Name("go.micro.web.heartbeat"),
	)

	if err := service.Init(); err != nil {
		log.Fatal("Init", err)
	}
	// websocket 连接接口 web.name注册根据.分割路由路径，所以注册的路径要和name对应上
	service.HandleFunc("/heartbeat", conn)

	if err := service.Run(); err != nil {
		log.Fatal("Run: ", err)
	}
}
func conn(w http.ResponseWriter, r *http.Request) {
	//
	c, err := upGrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("upgrade: %s", err)
		return
	}

	defer c.Close()

	for {
		mt, revMsg, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		sendMsg := MsgAssembler(revMsg)
		log.Printf("recv: %s", sendMsg)

		err = c.WriteMessage(mt, sendMsg)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func MsgAssembler(data []byte) []byte {
	// 转换为数字id
	b_buf := bytes.NewBuffer(data)
	var x int32
	binary.Read(b_buf, binary.BigEndian, &x)
	log.Printf("tick id = %d", x)
	return data
}
