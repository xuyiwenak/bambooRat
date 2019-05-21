package base

import (
	"auth_srv/base/config"
	"auth_srv/base/db"
	"auth_srv/base/redis"
)

func Init() {
	config.Init()
	db.Init()
	redis.Init()
}
