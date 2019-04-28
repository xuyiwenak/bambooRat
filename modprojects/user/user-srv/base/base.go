package base

import (
	"base/config"
	"base/db"
)

func Init() {
	config.Init()
	db.Init()
}
