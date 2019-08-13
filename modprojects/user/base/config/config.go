package config

import (
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source/consul"
	"github.com/micro/go-micro/util/log"
	"sync"
)

var (
	err error
)

var (
	defaultConfigPath       = "micro/config/cluster" // 默认的仓库地址
	defaultConsulServerAddr = "127.0.0.1:8500"
	consulConfig            defaultConsulConfig
	mysqlConfig             defaultMysqlConfig
	jwtConfig               defaultJwtConfig
	redisConfig             defaultRedisConfig
	m                       sync.RWMutex
	inited                  bool
)

// Init 初始化配置
func Init() {
	m.Lock()
	defer m.Unlock()
	if inited {
		log.Logf("[Init] 配置已经初始化过")
		return
	}
	// 从注册中心读取配置
	consulSource := consul.NewSource(
		consul.WithAddress(defaultConsulServerAddr),
		consul.WithPrefix(defaultConfigPath),
		consul.StripPrefix(true),
	)
	// 创建新的配置
	conf := config.NewConfig()
	if err := conf.Load(consulSource); err != nil {
		log.Logf("load config errr!!!", err)
	}
	// 侦听文件变动
	watcher, err := conf.Watch()
	if err != nil {
		log.Fatalf("[Init] 侦听consul配置中心 watcher异常，%s", err)
		panic(err)
	}
	go func() {
		for {
			v, err := watcher.Next()
			if err != nil {
				log.Fatalf("侦听consul配置中心 异常， %s", err)
				return
			}
			if err = conf.Load(consulSource); err != nil {
				panic(err)
			}
			log.Logf("consul配置中心有变化，%s", string(v.Bytes()))
		}
	}()
	// 赋值
	if err := conf.Get("consul").Scan(&consulConfig); err != nil {
		log.Logf("consul配置加载异常:%s", err)
	}
	if err := conf.Get("mysql").Scan(&mysqlConfig); err != nil {
		log.Logf("mysql配置加载异常:%s", err)
	}
	if err := conf.Get("redis").Scan(&redisConfig); err != nil {
		log.Logf("redis配置加载异常:%s", err)
	}
	if err := conf.Get("jwt").Scan(&jwtConfig); err != nil {
		log.Logf("jwt配置加载异常:%s", err)
	}

	// 标记已经初始化
	inited = true
}

// GetMysqlConfig 获取mysql配置
func GetMysqlConfig() (ret MysqlConfig) {
	return mysqlConfig
}

// GetConsulConfig 获取Consul配置
func GetConsulConfig() (ret ConsulConfig) {
	return consulConfig
}

// GetJwtConfig 获取Jwt配置
func GetJwtConfig() (ret JwtConfig) {
	return jwtConfig
}

// GetRedisConfig 获取Redis配置
func GetRedisConfig() (ret RedisConfig) {
	return redisConfig
}
