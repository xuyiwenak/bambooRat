# proto 项目  
## 目录结构  
```
├── README.md
├── auth
│   ├── auth.micro.go
│   ├── auth.pb.go
│   └── auth.proto
├── go.mod
├── go.sum
└── user
    ├── user.micro.go
    ├── user.pb.go
    └── user.proto
```  
proto 独立出来是为了方便其他项目交叉调用，同时可以统一生成对应的micro go文件。
执行脚本[proto生成脚本]
```
sh proto_gen_recurse.sh proto
```

[proto生成脚本]: https://github.com/xuyiwenak/bambooRat/blob/master/modprojects/user/proto_gen_recurse.sh
