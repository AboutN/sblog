package models

import (
	"sblog/log"
)

// Profile 博主的基本能信息
type Profile struct {
	Email    string `gorm:"primaryKey"` // 邮箱, 登录ID
	NickName string // 昵称
	Password string // 登录密码
	Address  string // 地址
	WeChat   string // 微信
	QQ       string // QQ
	Gitee    string // gitee帐号
	GitHub   string // github帐号
	Portrait string // base64 的头像
	Bio      string // 个人简介,
}

// Create 创建
func (m *Profile) Create() (result bool, err error) {
	err = conn.Create(m).Error
	if err != nil {
		result = false
		log.Error().Msg(err.Error())
		return
	}
	result = true
	return
}

// Get 获取基本信息
func (m *Profile) Get() (result *Profile, err error) {
	err = conn.First(result).Error
	if err != nil {
		result = nil
		log.Error().Msg(err.Error())
		return
	}
	return
}

// Edit 修改基本信息
func (m *Profile) Edit(param Params) (result bool) {
	err := conn.Updates(param).Error
	if err != nil {
		result = false
		log.Error().Msg(err.Error())
		return
	}
	result = true
	return
}
