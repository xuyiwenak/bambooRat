# base 项目
## 目录结构
基于go-micro写一个基本的发布/订阅例子 
- client  
  通常client作为发布端 pub
- proto  
  触发事件的协议载体
- server  
  通常server作为订阅端 sub

## 使用
micro 默认的pub、sub实现使用http，如果要更换broker为类似nats这样的插件在启动的时候需要指定。
### 普通启动方法
```
// 启动client 发布端
go run client.go
// 启动server 订阅端
go run server.go
```
### 使用nats启动方法

步骤：
1. 下载nats，推荐用docker环境装，比较简单
```
docker pull nats:latest
```
2. 在后台启动nats
```
docker run -d --name nats-main -p 4222:4222 -p 6222:6222 -p 8222:8222 nats
```
3. 启动服务
```
// 启动server
go run server.go --broker=nats --broker_address=127.0.0.1:4222
// 启动client
go run client.go --broker=nats --broker_address=127.0.0.1:4222
```