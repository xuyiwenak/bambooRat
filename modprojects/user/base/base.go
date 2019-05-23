package base

import (
	"github.com/xuyiwenak/bambooRat/modprojects/user/base/config"
	"github.com/xuyiwenak/bambooRat/modprojects/user/base/db"
	"github.com/xuyiwenak/bambooRat/modprojects/user/base/redis"
)

func Init() {
	config.Init()
	db.Init()
	redis.Init()
}
