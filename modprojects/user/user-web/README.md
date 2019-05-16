# User WEB Service

This is the User service

Generated with

```
micro new user-web --namespace=bambooRat.micro.sdk --alias=user --type=web --gopath=false
```

## Getting Started

- [配置环境](#configuration)
- [依赖](#dependencies)
- [测试](#test)
- [具体使用](#usage)


## Configuration

- FQDN: bambooRat.micro.web.user
- Type: web
- Alias: user

## Dependencies
### consul
推荐用consul进行服务发现检测，如果没有安装consul可以执行[tools/framework_install.sh][framework_install.sh]

```
# mac install consul
brew install consul

# run consul
consul agent -dev
```
查看进程状态  
http://localhost:8500/ui/dc1/services
## Usage

主要是以docker环境为主，为了方便docker获取编译时需要的依赖库，需要在当前项目下简历vendor目录，执行
```
# go module 本身是支持vendor模式的，把需要的代码拷贝进来,build会从vendor目录获取代码
go mod vendor
```
制作docker镜像
```
docker build -t user-web .
```
运行docker镜像
```
# 养成加--rm的好习惯，不然container会越来越多
docker run -it --rm --name user-web user-web
```

[framework_install.sh]: https://github.com/xuyiwenak/bambooRat/blob/master/tools/framework_install.sh