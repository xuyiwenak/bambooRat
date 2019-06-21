# base 项目  
## 目录结构  
```
.
├── README.md
├── base.go
├── config
│   ├── config.go
│   ├── consul.go
│   ├── jwt.go
│   ├── mysql.go
│   ├── profiles.go
│   └── redis.go
├── db
│   ├── db.go
│   └── mysql.go
├── go.mod
├── go.sum
└── redis
    └── redis.go
```  
主要存放基础组件模块，用来处理第三方调用，同时其他模块可以通过调用base实现复用。

如果要使用docker-compose启动需要放置在[modproject/user][user]项目下  

[user]:https://github.com/xuyiwenak/bambooRat/tree/master/modprojects/user