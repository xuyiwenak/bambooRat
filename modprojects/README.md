# go micro基础结构
## micro main function 的基本编写方法
1. 构造micro
2. 初始化micro命令行
3. 注册
4. 运行

示例代码：
```
package main

import (
	"context"

	proto "github.com/micro/examples/function/proto"
	"github.com/micro/go-micro"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}

func main() {
	// 创建新函数
	fnc := micro.NewFunction(
		micro.Name("greeter"),
	)

	// 初始化命令行
	fnc.Init()

	// 注册handler
	fnc.Handle(new(Greeter))

	// 运行服务
	fnc.Run()
}
```