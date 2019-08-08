package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	"github.com/pborman/uuid"
	"time"
	proto "web/pubsub/base/proto"
)

// sendEv 通过publisher发布事件
func sendEv(topic string, p micro.Publisher) {
	t := time.NewTicker(time.Second)
	counter := 0
	for _ = range t.C {
		counter++
		// 创建一个新的事件
		ev := &proto.Event{
			Id:        uuid.NewUUID().String(),
			Timestamp: time.Now().Unix(),
			Message:   fmt.Sprintf("Messaging send env:%s number:%d", topic, counter),
		}
		log.Logf("publishing %+v\n", ev)
		// 发布一个事件
		if err := p.Publish(context.Background(), ev); err != nil {
			log.Logf("error publishing: %v", err)
		}
	}
}

func main() {
	// 创建一个micro服务
	service := micro.NewService(
		micro.Name("bambooRat.micro.cli.pubsub"),
	)
	// 初始化服务
	service.Init()

	// 创建发布者
	pub1 := micro.NewPublisher("pubsub1", service.Client())
	pub2 := micro.NewPublisher("pubsub2", service.Client())

	// pub to topic 1
	go sendEv("bambooRat.topic.pubsub.1", pub1)
	// pub to topic 2
	go sendEv("bambooRat.topic.pubsub.2", pub2)

	// 永远阻塞
	select {}
}
