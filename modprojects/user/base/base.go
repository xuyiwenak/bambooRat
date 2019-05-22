package base

import (
	"user/base/config"
	"user/base/db"
	"user/base/redis"
)

func Init() {
	config.Init()
	db.Init()
	redis.Init()
}
