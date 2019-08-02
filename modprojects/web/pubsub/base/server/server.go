package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-micro/util/log"
	proto "web/pubsub/base/proto"

	"context"
)

// 当Sub 被消息触发执行的时候，所有的方法都会执行
type Sub struct{}

// 这个方法可以是任何名字
func (s *Sub) Process(ctx context.Context, event *proto.Event) error {
	md, _ := metadata.FromContext(ctx)
	log.Logf("[pubsub1] Received event %+v with metadata %+v\n", event, md)
	// do something with event
	return nil
}

// 自定义的方法
func subEv(ctx context.Context, event *proto.Event) error {
	md, _ := metadata.FromContext(ctx)
	log.Logf("[pubsub2] Received event %+v with metadata %+v\n", event, md)
	// do something with event
	return nil
}

func main() {
	// 创建一个micro服务
	service := micro.NewService(
		micro.Name("bambooRat.micro.srv.pubsub"),
	)
	// 初始化服务
	service.Init()

	// 注册一个订阅者
	if err := micro.RegisterSubscriber("pubsub1", service.Server(), new(Sub)); err != nil {
		log.Fatal(err)
	}
	// 通过队列注册一个订阅者，每条消息都会分发给第一无二的订阅者
	if err := micro.RegisterSubscriber("pubsub2", service.Server(), subEv, server.SubscriberQueue("queue.pubsub")); err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
