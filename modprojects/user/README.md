# user项目说明 
##  项目结构
```
├── README.md
├── base                 # 基础组件库
│   ├── README.md
│   ├── base.go
│   ├── config
│   ├── db
│   ├── go.mod
│   ├── go.sum
│   └── redis
├── docker-compose.yml   # docker-compose 工具库，使用docker容器编排需要注意修改业务conf里面的连接名称
├── proto                # proto类库
│   ├── README.md
│   ├── auth
│   ├── go.mod
│   ├── go.sum
│   └── user
├── proto_gen_recurse.sh # 生成proto类文件的脚本
├── scripts
│   └── schema.sql
├── user-auth            # bambooRat.micro.srv.auth 处理认证业务
│   ├── Dockerfile
│   ├── README.md
│   ├── conf             # 业务的配置库
│   ├── go.mod
│   ├── go.sum
│   ├── handler
│   ├── main.go
│   ├── model
│   ├── plugin.go
│   └── vendor
├── user-srv            # bambooRat.micro.srv.user 对架构内应用提供业务服务
│   ├── Dockerfile
│   ├── README.md
│   ├── conf
│   ├── go.mod
│   ├── go.sum
│   ├── handler
│   ├── main.go
│   ├── model
│   ├── plugin.go
│   └── vendor
└── user-web            # bambooRat.micro.web.user 接收API下放的路由请求
    ├── Dockerfile
    ├── README.md
    ├── conf
    ├── get_url_test.sh
    ├── go.mod
    ├── go.sum
    ├── handler
    ├── main.go
    ├── plugin.go
    └── vendor
``` 

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
