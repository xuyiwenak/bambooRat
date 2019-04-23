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

This directory contains a grpc gateway generated using [grpc-ecosystem/grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway).

Services written with [micro/go-grpc](https://github.com/micro/go-grpc) are fully compatible with the grpc-gateway and any other 
grpc services.

Go to [grpc-ecosystem/grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway) for details on how to generate gateways. We 
have generated the gateway from the same proto as the greeter server but with additional options for the gateway.

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

