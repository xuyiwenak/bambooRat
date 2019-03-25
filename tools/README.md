# 1. 启动consul代理  
-  consul 集群启动  
```
consul agent -dev -node machine  
```
启动本地的节点machine  
-  查看成员  
```
consul members
```
-  http协议访问节点    
```
curl localhost:8500/v1/catalog/nodes
```
# 2. 注册服务
- 定义一个服务
可以通过提供服务定义或通过对HTTP API进行适当调用来注册服务  

```  
mkdir -p ${GOPATH}/confs/consul.d
cd ${GOPATH}/confs/consul.d
echo '{"service": {"name": "web", "tags": ["rails"], "port": 80}}'>web.json
```  
- 查询服务  
启动代理并同步服务后，我们可以使用HTTP API查询服务  
```
curl http://localhost:8500/v1/catalog/service/web
```
- API相关文档  
https://www.consul.io/api/agent/service.html
