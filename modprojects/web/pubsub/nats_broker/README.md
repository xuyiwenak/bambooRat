# nats_broker
一个基于nats_broker 的发布订阅模板
## 需求
- nats
## 步骤
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
参考自[Bruce Wang][https://github.com/BruceWangNo1/go-micro-pubsub-with-nats]






