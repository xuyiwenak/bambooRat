package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/broker/nats"
	"github.com/micro/go-micro/util/log"
	"github.com/pborman/uuid"
	"time"
	proto "web/pubsub/base/proto"
)

func SendEv(topic string, p micro.Publisher) {
	t := time.NewTicker(5 * time.Second)

	counter := 0
	for _ = range t.C {
		counter++
		ev := proto.Event{
			Id:        uuid.NewUUID().String(),
			Timestamp: time.Now().Unix(),
			Message:   fmt.Sprintf("pub topic name %s", topic),
		}

		body, _ := json.Marshal(ev)
		msg := &broker.Message{
			Header: map[string]string{
				"id": fmt.Sprintf("%d", counter),
			},
			Body: body,
		}

		log.Logf("publishing %+v", ev)

		if err := p.Publish(context.Background(), msg); err != nil {
			log.Logf("error publishing: %v", err)
		}

	}
}

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("bambooRat.micro.cli.pub"),
		micro.Version("latest"),
		micro.Broker(nats.NewBroker()),
	)
	pubTopic1 := "bambooRat.micro.pubsub.topic.event"
	pub1 := micro.NewPublisher(pubTopic1, service.Client())
	// service初始化会一并初始化broker 的init
	service.Init()

	go SendEv(pubTopic1, pub1)
	// 启动service的时候执行broker的connect
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
