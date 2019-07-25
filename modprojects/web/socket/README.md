# websocket 项目

```
├── go.mod
├── go.sum
├── socket
│   ├── README.md
│   ├── heartbeat_demo      // 客户端到服务器的心跳模型
│   ├── service_demo        // 客户端手动发送终端输入的自定义消息到服务的模型
│   └── web_demo            // web页面发送消息到服务器的模型
```
该项目都是基于micro api web 所以需要先启动web服务
```
micro api --handler=web --namespace=go.micro.web
```
socket module for test
