# User Web 服务

这是处理user的web部分

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
### docker-compose
一般docker环境完成后就可以使用docker-compose，添加对应的docker-compose配置即可
## Usage
### docker
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
### docker-compose(简单，推荐)
如果是基于本地测试集群推荐docker-compose，容器会自动根据名称连接，而不是host固定的ip端口
```
# 启动
docker-compose up -d 
```
会自动根据docker-compose.yml配置的镜像启动对应的容器，如果没有可以自动下载

```
# 停止，自动移除镜像
docker-compose down
```  
### 调试
如果是需要反复调试代码，可以修改conf 配置切换成本地host，再编译执行
```
# 编译
go build -o xxxx main.go xxxx.go
# 运行
./xxxx
```


[framework_install.sh]: https://github.com/xuyiwenak/bambooRat/blob/master/tools/framework_install.sh