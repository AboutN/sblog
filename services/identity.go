package services

import (
	"sblog/entities"
	"sblog/models"
	"sblog/scode"
	"sblog/utils"
)

//
func ManageService(userid, password string) entities.ResultEntity {
	model := &models.Profile{
		Email: userid,
	}
	result, err := model.Get()
	if err != nil {
		return entities.ResultEntity{
			Code: scode.ErrorDB,
			Info: "",
		}
	}
	if result.Password == utils.MD5(password) && result.Email == userid {
		return entities.ResultEntity{
			Code: scode.Success,
			Info: "登录成功",
		}
	}

	return entities.ResultEntity{
		Code: scode.Success,
		Info: "用户名或密码错误",
	}
}
