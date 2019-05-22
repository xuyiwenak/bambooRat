# User Service

This is the User service

Generated with

```
micro new user-srv --namespace=bambooRat.micro.sdk --alias=user --type=srv --gopath=false
```

## Getting Started

- [配置环境](#configuration)
- [依赖](#dependencies)
- [测试](#test)
- [具体使用](#usage)


## Configuration

- FQDN: bambooRat.micro.srv.user
- Type: srv
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
### mysql
推荐直接docker 拉取镜像安装自己需求的版本
```
# docker 中下载 mysql 现在默认是8.x 如果想下载5.x可以显式指定mysql5.6
docker pull mysql

#启动 启动的端口和密钥自行设置
docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 -d mysql

#进入容器
docker exec -it mysql bash

#登录mysql
mysql -u root -p
ALTER USER 'root'@'localhost' IDENTIFIED BY '123456';

#添加远程登录用户
CREATE USER 'guest_user'@'%' IDENTIFIED WITH mysql_native_password BY '123456';
GRANT ALL PRIVILEGES ON *.* TO 'guest_user'@'%';
```
## test
```
# 启动consul 打开监听
$ consul agent -dev
# 执行
$ go run main.go plugin.go
```  
查看consul进程状态  
http://localhost:8500/ui/dc1/services
## Usage

主要是以docker环境为主，为了方便docker获取编译时需要的依赖库，需要在当前项目下简历vendor目录，执行
```
# go module 本身是支持vendor模式的，把需要的代码拷贝进来
go mod vendor
```
制作docker镜像
```
docker build -t user-srv .
```
运行docker镜像
```
# 养成加--rm的好习惯，不然container会越来越多
docker run -it --rm --name user-srv user-srv
```

[framework_install.sh]: https://github.com/xuyiwenak/bambooRat/blob/master/tools/framework_install.sh