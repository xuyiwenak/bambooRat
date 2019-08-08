package main

import (
	"context"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker/nats"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/util/log"
	proto "web/pubsub/base/proto"
)

// 当Sub 被消息触发执行的时候，所有的方法都会执行
type Sub struct{}

func main() {
	// 初始化之前指定nats作为broker
	service := micro.NewService(
		micro.Name("bambooRat.micro.srv.pubsub"),
		micro.Version("latest"),
		micro.Broker(nats.NewBroker()),
	)
	// 注册一个订阅者
	if err := micro.RegisterSubscriber("bambooRat.micro.pubsub.topic.event", service.Server(), new(Sub)); err != nil {
		log.Fatal(err)
	}
	// service初始化会一并初始化broker 的init
	service.Init()
	// 启动service的时候执行broker的connect
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
func (s *Sub) Process(ctx context.Context, event *proto.Event) error {
	md, _ := metadata.FromContext(ctx)
	log.Logf("[sub] Received event %+v with metadata %+v\n", event, md)
	// do something with event
	return nil
}
