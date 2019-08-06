package main

import (
	"encoding/json"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/util/log"
	"github.com/pborman/uuid"
	"time"
	proto "web/pubsub/base/proto"
)

func SendEv(topic string, b broker.Broker) {
	t := time.NewTicker(5 * time.Second)

	var i int
	for _ = range t.C {
		ev := proto.Event{
			Id:        uuid.NewUUID().String(),
			Timestamp: time.Now().Unix(),
			Message:   fmt.Sprintf("Message you all day on %s, 'cause nothing's gonna change my love for you", topic),
		}

		body, _ := json.Marshal(ev)
		msg := &broker.Message{
			Header: map[string]string{
				"id": fmt.Sprintf("%d", i),
			},
			Body: body,
		}

		log.Logf("publishing %+v", ev)

		if err := b.Publish(topic, msg); err != nil {
			log.Logf("error publishing: %v", err)
		}
		i++
	}
}

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.cli.pubsub"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	b := service.Client().Options().Broker

	if err := b.Init(); err != nil {
		log.Fatalf("Broker Init error: %v", err)
	}
	if err := b.Connect(); err != nil {
		log.Fatalf("Broker Connect error: %v", err)
	}

	go SendEv("go.micro.pubsub.topic.event", b)
	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
