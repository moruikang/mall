package service

import (
	"github.com/pkg/errors"
	"mall.com/config/global"
	"mall.com/store/models"
)

type WebUserService struct {
}

// Login 用户登录信息验证
func (u *WebUserService) Login(param models.WebUserLoginParam) uint64 {
	var user models.User
	global.Db.Where("username = ? and password = ?", param.Username, param.Password).First(&user)
	return user.Id
}

func GetUser(param models.WebUserLoginParam) (userinfo models.User, err error) {
	var user models.User
	global.Db.Where("username = ? and password = ?", param.Username, param.Password).First(&user)

	if user.Id > 0 {
		return user, nil
	}
	return user, errors.New("用户不存在")
}
