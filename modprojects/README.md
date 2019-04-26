# micro 项目
## 自动构建micro模板项目
micro new的使用方法,查看[源码][micro/new源码]获取详细信息.
```
micro new -h
NAME:
   micro new - 创建一个micro模板

USAGE:
   micro new [command options] [arguments...]

OPTIONS:
   --namespace "go.micro"                       执行micro服务的名空间
   --type "srv"                                 该micro服务的类型，支持四种 api、fnc(function)、srv(service)、web。
   --fqdn                                       默认情况下，服务全名就是命名空间+服务类型+服务名，修改会替换掉 e.g com.example.srv.service (defaults to namespace.type.alias)
   --alias                                      一般最后一个名字就是服务名，通过参数可以替换 e.g. micro  --> orcim
   --plugin [--plugin option --plugin option]   指定registry和broker所使用的第三方插件 --plugin=registry=etcd:broker=nats or use flag multiple times
   --gopath                                     默认是使用GOPATH，会在${GOPTAH}/src/下生成对应的项目。如果使用go module，需要  
                                                设置为false，在指定的目录直接生成
```
例如生成一个task的例子
```
$ micro new --gopath=false task

Creating service go.micro.srv.task in task

.
├── main.go
├── plugin.go
├── handler
│   └── example.go
├── subscriber
│   └── example.go
├── proto/example
│   └── example.proto
├── Dockerfile
├── Makefile
├── README.md
└── go.mod


download protobuf for micro:

brew install protobuf
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
go get -u github.com/micro/protoc-gen-micro

compile the proto file example.proto:

cd task
protoc --proto_path=. --go_out=. --micro_out=. proto/example/example.proto
```
tips: 如果proto文件有其他proto的依赖关系，需要在--proto_path里面添加新的路径，而且需要保证待生成的.proto  
      文件也在--proto_path的包括中


## go micro基础结构
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

[micro/new源码]: https://github.com/micro/micro/blob/master/new/new.go