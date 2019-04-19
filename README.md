# bambooRat
***
## 项目简介
主要是方便自动构建项目，快速进行基础开发  
本项目提供了两种构建方式, 任意一种或者两种可以同时使用,后面会详细解释。   
* 依赖GOPATH 的传统构建方式，在目录goproject下
* 不依赖GOPATH，使用go module管理依赖包，并且可以和GOPATH同步使用的自动构建方式(推荐)  
```
├── LICENSE 
├── README.md
├── go  // golang bin dir
├── goprojects // GOPATH project
├── modprojects // mod project
├── start.sh // init script
└── tools // constract scripts
```
### 1. goprojects  
sh start.sh 
一键安装， golang语言包，golang环境变量， protobuf micro grpc等基本安装包
构造项目的环境变量
这里被墙掉的库需要在tools下面的github_list文本进行手动维护，可以仿照格式自行添加。
### 2. modprojects  
```
cd modprojects
go mod tidy // download dependances
go bulid ${project_name} // your project name
```  
被墙掉的url自己redirect到github对应的目录上，初始的版本已经比较详细，如果出现部分超时的情况，可以到go.mod文件里自行添加。  
后面也会定期维护。  
ps: go dep 虽然目前官方也有支持，需要搭建私有仓库。这里主要基于公共框架，所以没有列入该方案。  
dep在处理被墙的域名方面没有优势，目前官方从1.11版本以后主要推荐go module，如果自行开发推荐使用。
相关私有仓库的搭建推荐使用[gogs](https://github.com/gogs/gogs)  
如果不想在获取依赖包时太复杂，也可以使用第三方提供的包管理工具[gopm](https://github.com/gpmgo/gopm)  
缺点就是要依赖gopm官方维护，经常有小问题。



