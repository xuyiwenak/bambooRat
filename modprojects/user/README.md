# user项目说明 

- user-web 以下简称**web**
- user-srv 以下简称**service**

|服务|命名空间|说明|---|
|---|---|---|---|
|接入层API|bambooRat.micro.sdk.web|负责代理所有**bambooRat.micro.sdk.web**下游的web应用，比如**bambooRat.micro.sdk.web.user**等|---|
|用户web|bambooRat.micro.sdk.web.user|接收API下放的路由为/user请求|---|
|用户服务|bambooRat.micro.sdk.srv.user|对架构内应用提供user查询服务|---|  
 
```
micro new user-srv --namespace=bambooRat.micro.sdk --alias=user --type=srv --gopath=false
```  
docker 启动命令
```
docker run -p 3306:3306 --name mysql -v $HOME/mysql/conf:/etc/mysql/conf.d -v $HOME/mysql/logs:/logs -v $HOME/mysql/data:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=123456 -d mysql
```
