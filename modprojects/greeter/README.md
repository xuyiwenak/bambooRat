# Greeter

An example Greeter application
一个Greeter的应用

## Contents 目录

- **srv** - an RPC greeter service - RPC服务的位置  
- **cli** - an RPC client that calls the service once - 调用的客户端服务，这个客户端只会调用一次 
- **api** - examples of RPC API and RESTful API - 调用接口  
- **web** - how to use go-web to write web services - go-web服务相关

## Run Service  启动micro服务

Start go.micro.srv.greeter
```shell
go run srv/main.go
```

## Client 启动客户端

Call go.micro.srv.greeter via client
```shell
go run cli/main.go
```

Examples of client usage via other languages can be found in the client directory.
文件夹内有其他语言的调用方式

## API

HTTP based requests can be made via the micro API. Micro logically separates API services from backend services. By default the micro API 
accepts HTTP requests and converts to *api.Request and *api.Response types. Find them here [micro/api/proto](https://github.com/micro/micro/tree/master/api/proto).
可以通过micro 的接口 生成基于HTTP的请求。  

Run the go.micro.api.greeter API Service 启动micro的greeter api
```shell
go run api/api.go 
```

Run the micro API
```shell
micro api --handler=api
```

Call go.micro.api.greeter via API
```shell
curl http://localhost:8080/greeter/say/hello?name=John
```

Examples of other API handlers can be found in the API directory.

## Web
micro web是一个记录web状态的仪表盘， 或者作为代理运行micro的服务  
The micro web is a web dashboard and reverse proxy to run web apps as microservices.  

Run go.micro.web.greeter
```
go run web/web.go 
```

Run the micro web
```shell
micro web
```

Browse to http://localhost:8082/greeter  
