package config

import (
	"fmt"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source/consul"
	"github.com/micro/go-micro/util/log"
	"strings"
	"sync"
)

var (
	err error
)

var (
	defaultConfigPath       = "/micro/config/cluster" // 默认的仓库地址
	defaultConsulServerAddr = "127.0.0.1:8500"
	consulConfig            defaultConsulConfig
	mysqlConfig             defaultMysqlConfig
	jwtConfig               defaultJwtConfig
	redisConfig             defaultRedisConfig
	profiles                defaultProfiles
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
		// optionally strip the provided prefix from the keys, defaults to false
		consul.StripPrefix(true),
	)
	// 创建新的配置
	conf := config.NewConfig()
	if err := conf.Load(consulSource); err != nil {
		log.Logf("load config errr!!!", err)
	}
	if err := conf.Get("micro", "config", "cluster"); err != nil {
		log.Logf("json format err!!!", err)
	}
	configMap := conf.Map()
	fmt.Println(configMap)

	// 侦听文件变动
	watcher, err := conf.Watch()
	if err != nil {
		log.Fatalf("[Init] 开始侦听应用配置文件变动 异常，%s", err)
		panic(err)
	}
	go func() {
		for {
			v, err := watcher.Next()
			if err != nil {
				log.Fatalf("[loadAndWatchConfigFile] 侦听应用配置文件变动 异常， %s", err)
				return
			}
			if err = conf.Load(consulSource); err != nil {
				panic(err)
			}
			log.Logf("[loadAndWatchConfigFile] 文件变动，%s", string(v.Bytes()))
		}
	}()
	// 处理前缀分割问题
	prefixPath := strings.Replace(defaultConfigPath, "/", ",", -1)
	// 赋值
	conf.Get(strings.TrimLeft(prefixPath, ","), "consul").Scan(&consulConfig)
	conf.Get(strings.TrimLeft(prefixPath, ","), "mysql").Scan(&mysqlConfig)
	conf.Get(strings.TrimLeft(prefixPath, ","), "redis").Scan(&redisConfig)
	conf.Get(strings.TrimLeft(prefixPath, ","), "jwt").Scan(&jwtConfig)

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
