package main

import (
	"github.com/gogo/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/micro/go-micro/web"
	"log"
	"net/http"
	"web/socket/heartbeat_demo/proto"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}
var (
	clientRes heartbeat.Request
)

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
	conn, err := upGrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("upgrade: %s", err)
		return
	}

	defer conn.Close()

	for {
		_, buffer, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		if err := proto.Unmarshal(buffer, &clientRes); err != nil {
			log.Printf("proto unmarshal: %s", err)
		}
		log.Printf("recv userId=%d MsgId=%d Data=%s", clientRes.UserId, clientRes.MsgId, clientRes.Data)
	}
}
