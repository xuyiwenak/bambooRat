package main

import (
	"context"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/errors"
	proto "github.com/xuyiwenak/bambooRat/modprojects/micro-api/RPC/proto"
)

type Alice struct {
}
type Bob struct {
}

func (pb *Alice) Sing(ctx context.Context, req *proto.AliceRequest, rsp *proto.AliceResponse) error {
	log.Logf("Alice.Sing rev request")

	var rev_msg string = req.RevMsg

	if len(rev_msg) == 0 {
		return errors.BadRequest("go.micro.rpc.Alice", "参数不正确")
	}
	rsp.SendMsg = "congratulations to" + rev_msg
	rsp.RetCode = 1
	return nil
}
func (pb *Bob) Dance(ctx context.Context, req *proto.BobRequest, rsp *proto.BobResponse) error {
	log.Logf("Bob.Dance rev request")

	var rev_msg string = req.RevMsg

	if len(rev_msg) == 0 {
		return errors.BadRequest("go.micro.rpc.Bob", "参数不正确")
	}
	rsp.SendMsg = "congratulations to" + rev_msg
	rsp.RetCode = 1
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.rpc.demo"),
	)

	service.Init()

	// 注册 example handler
	proto.RegisterAliceHandler(service.Server(), new(Alice))

	// 注册 foo handler
	proto.RegisterBobHandler(service.Server(), new(Bob))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
