# RPC API

本示例我们介绍如何使用RPC handler模式的**Micro API**，以下简称**API**。

该模式下允许我们通过RPC的方式把HTTP请求转发到go-micro微服务上。

需要提醒的是，RPC模式下**API**只接收POST方式的请求，并且只支付内容格式**content-type**为**application/json**或者**application/protobuf**。

## 使用方法

以rpc模式运行**API**

```
micro api --handler=rpc
```

```
go run rpc.go
```

当我们POST请求到 **demo/Alice/sing**时，**API**会将它转成RPC转发到**go.micro.api.demo**服务的**Alice.sing**接口上。 

```
curl -H 'Content-Type: application/json' -d '{"rev_msg": "evan"}' "http://localhost:8080/demo/Alice/sing"
```

同样，POST请求到 **demo/Bob/dance**时，**API**会将它转成RPC转发到**go.micro.api.demo**服务的**Bob.dance**接口上。

```
curl -H 'Content-Type: application/json' -d '{"rev_msg": "evan"}' "http://localhost:8080/demo/Bob/dance"
```
