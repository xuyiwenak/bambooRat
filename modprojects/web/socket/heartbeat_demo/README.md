# heartbeat_demo 项目
## server
基于micro web api的服务，需要启动web api
1. 启动micro web api
```
// 启动micro web api
micro api --handler=web --namespace=go.micro.web
```
2. 启动websocket server
```
// 启动websocket server
go run server/server.go
```
## client
1. 启动client发送消息
基于micro service的客户端，发起websocket连接。
```
// 启动客户端
go run client/client.go

```



