package entities

import (
	"sblog/scode"
)

// ResultEntity 请求结果实体
type ResultEntity struct {
	Code scode.StatusCode `json:"code"` // 状态码
	Info string           `json:"info"` // 信息
	Data interface{}      `json:"data"` // 数据
}

// ManageLogin  博主登录
type ManageLogin struct {
	UserID   string `json:"userid"`   // 登录ID Email
	Password string `json:"password"` //
}
