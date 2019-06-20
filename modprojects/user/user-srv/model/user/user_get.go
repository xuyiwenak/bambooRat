package user

import (
	"github.com/micro/go-micro/util/log"
	"github.com/xuyiwenak/bambooRat/modprojects/user/base/db"
	proto "github.com/xuyiwenak/bambooRat/modprojects/user/proto/user"
)

func (s *service) QueryUserByName(userName string) (ret *proto.User, err error) {

	queryString := `SELECT user_id, user_name, pwd FROM user WHERE user_name = ?`

	// 获取数据库
	o := db.GetDB()

	ret = &proto.User{}

	// 查询
	err = o.QueryRow(queryString, userName).Scan(&ret.Id, &ret.Name, &ret.Pwd)
	if err != nil {
		log.Logf("[QueryUserByName] 查询数据失败，err：%s", err)
		return
	}
	return
}
