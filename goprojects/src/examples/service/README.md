# Service

这是创建一个 micro service 的例子.

## 目录

- main.go - 服务的定义, 包括handler和客户端
- proto - protobuf 的定义文件
- gen_micro_proto.sh - 根据proto目录下存放的.proto文件生成对应的go-micro文件 

## 运行

Run the service

```shell
go run main.go
```

Run the client

```shell
go run main.go --run_client
```
输出 Hello John