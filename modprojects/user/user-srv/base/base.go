package base

import (
	"github.com/xuyiwenak/bambooRat/modprojects/user/user-srv/base/config"
	"github.com/xuyiwenak/bambooRat/modprojects/user/user-srv/base/db"
)

func Init() {
	config.Init()
	db.Init()
}
