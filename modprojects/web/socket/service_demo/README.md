# service_demo 项目
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
// 输入发送到server的消息,如下
2019/07/25 19:58:27 Transport [http] Listening on [::]:63074
2019/07/25 19:58:27 Broker [http] Connected to [::]:63076
2019/07/25 19:58:27 Registry [mdns] Registering node: go.micro.web.wsClient-2d95c380-71cc-4d38-b331-8a610d826d86
2019/07/25 19:58:27 input your message here:%!(EXTRA <nil>)
1
2019/07/25 19:59:08 send msg client to server values:==>1
2019/07/25 19:59:08 recv msg from server values: 1<===server
2
2019/07/25 19:59:11 send msg client to server values:==>2
2019/07/25 19:59:11 recv msg from server values: 2<===server

```



