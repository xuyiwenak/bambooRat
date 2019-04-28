package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"user-srv/handler"
	user "user-srv/proto/user"
	"user-srv/subscriber"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("bambooRat.micro.sdk.srv.user"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	user.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
