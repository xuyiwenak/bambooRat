# GRPC

Contains examples for using [go-grpc](https://github.com/micro/go-grpc)

- [greeter](greeter) - A greeter example
- [gateway](gateway) - A grpc gateway example

## New service

Check out the [greeter](greeter) example using go-grpc

### Import go-grpc

```
import "github.com/micro/go-grpc"
```

### Create micro.Service

```
service := grpc.NewService()
```

## Pre-existing Service

What if you want to add grpc to a pre-existing service? Use the build pattern for plugins but swap out the client/server.

### Create a plugin file

```
package main

import (
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/server"
	cli "github.com/micro/go-plugins/client/grpc"
	srv "github.com/micro/go-plugins/server/grpc"
)

func init() {
	// set the default client
	client.DefaultClient = cli.NewClient()
	// set the default server
	server.DefaultServer = srv.NewServer()
}
```

### Build the binary

```
// For local use
go build -i -o service ./main.go ./plugins.go
```

### Run

Because the default client/server have been replaced we can just run as usual

```
./service

```

# GRPC Gateway  
gateway目录主要处理url访问的请求，通过grpc转发给对应的业务greeter，并返回相对应的结果。
![grpc-gateway的架构介绍](https://docs.google.com/drawings/d/12hp4CPqrNPFhattL_cIoJptFvlAqm5wLQ0ggqI5mkCg/pub?w=749&amp;h=370)
项目代码参考[grpc-ecosystem/grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway).    
这部分依赖三个库
- protoc-gen-grpc-gateway  
- protoc-gen-grpc-swagger  
- protoc-gen-go

这三个目录在[tools/framework_install.sh](https://github.com/xuyiwenak/bambooRat/tree/master/tools/framework_install.sh), 
该脚本已经安装过，可以前往GOBIN目录查看，确定安装成功后  


Services written with [micro/go-grpc](https://github.com/micro/go-grpc) are fully compatible with the grpc-gateway and any other 
grpc services.


## Usage

Run the go.micro.srv.greeter service

```
go run ../greeter/srv/main.go --server_address=localhost:9090
```

Run the gateway

```
go run main.go
```

Curl your request at the gateway (localhost:8080)

```
curl -d '{"name": "evan"}' http://localhost:8080/greeter/hello
```

