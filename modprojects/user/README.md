# user项目说明 
## 项目目录结构说明
- user-web 以下简称**web**
- user-srv 以下简称**service**

|服务|命名空间|说明|---|
|---|---|---|---|
|接入层API|bambooRat.micro.sdk.web|负责代理所有**bambooRat.micro.sdk.web**下游的web应用，比如**bambooRat.micro.sdk.web.user**等|---|
|用户web|bambooRat.micro.sdk.web.user|接收API下放的路由为/user请求|---|
|用户服务|bambooRat.micro.sdk.srv.user|对架构内应用提供user查询服务|---|  
 
 ##docker构建微服务方式
 为了方便微服务的管理，单机环境目前由docker-compose做容器编排，降低开发成本，具体步骤如下：
 - 启动
 ```
 docker-compose up -d
 # 启动完成后查看容器启动状态
 docker ps -a 
 ```
 这里有一点需要注意，虽然docker-compose提供了depends依赖关系，但是像mysql这类的数据库启动较慢的容器还是无法
 保证启动的先后顺序。 可以通过增加脚本对端口监听重新拉起服务。我目前的方案比较简单，如果服务对基础镜像强依赖关系
 则等依赖的容器启动后，再手动重启对应的服务即可。
 ```
 docker-compose restart user-srv
 # 或者重新执行
  docker-compose up -d
 ```
 如果需要关闭所有服务执行
 ```
 docker-compose down
 ```
 ## 构建mysql表
 把sql脚本拷贝到容器内的某个文件下，以home为例
  ```
  docker cp scripts/schema.sql <container name>:/home
  ```
 进入容器,执行脚本,创建数据库和表
 ```
 docker exec -it <container name> bash
 mysql -uroot -p -D micro_user</home/schema.sql
 ```
 设置数据库权限配置
 ```
 # 先进入数据库
 mysql -uroot -p
 # 执行sql语句
select user,host,authentication_string from mysql.user; 
 ```
 为root用户分配远程权限
 ```
 grant all PRIVILEGES on *.* to root@'%' WITH GRANT OPTION;
 ```
 由于Mysql5.6以上的版本修改了Password算法，这里还需要更新密码算法，方便远程连接
 ```
 ALTER user 'root'@'%' IDENTIFIED BY '123456' PASSWORD EXPIRE NEVER;
 ALTER user 'root'@'%' IDENTIFIED WITH mysql_native_password BY '123456';
 FLUSH PRIVILEGES;
 ```
