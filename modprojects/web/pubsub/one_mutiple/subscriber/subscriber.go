package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/util/log"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("bambooRat.micro.srv.billboard"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	broker := service.Server().Options().Broker

	if err := broker.Init(); err != nil {
		log.Fatalf("Broker Init error: %v", err)
	}
	if err := broker.Connect(); err != nil {
		log.Fatalf("Broker Connect error: %v", err)
	}

	if _, err := broker.Subscribe("topic1", Handler1); err != nil {
		log.Fatalf("broker.Subscribe error: %v", err)
	}
	if _, err := broker.Subscribe("topic2", Handler); err != nil {
		log.Fatalf("broker.Subscribe error: %v", err)
	}
	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
func Handler1(event broker.Event) error {
	log.Logf("[sub] topic1: %v, %v", event.Message().Header, string(event.Message().Body))
	return nil
}
func Handler(event broker.Event) error {
	log.Logf("[sub] topic2: %v, %v", event.Message().Header, string(event.Message().Body))
	return nil
}
