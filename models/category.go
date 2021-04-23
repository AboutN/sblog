package models

import "sblog/log"

//  Category 文章类别
type Category struct {
	ID          string // 编号
	CName       string `gorm:"column:cname"` // 分类名称
	Description string // 分类概述
}

// Category 创建分类
func (m *Category) Create() (result bool) {
	err := conn.Create(m).Error
	if err != nil {
		result = false
		log.Error().Msg(err.Error())
		return
	}
	result = true
	return
}

// Get 根据分类名称获取
func (m *Category) Get(cname string) (result *Category) {
	db := conn.Where("cname = ?", cname).First(result)
	err := db.Error
	if err != nil || db.RowsAffected == 0 {
		result = nil
		log.Error().Msg(err.Error())
		return
	}
	return
}

// Edit 修改分类
func (m *Category) Edit(param Params) (result bool) {
	err := conn.Updates(param).Error
	if err != nil {
		result = false
		log.Error().Msg(err.Error())
		return
	}
	result = true
	return
}

// QueryAll 查询所有的分类
func (m *Category) QueryAll() (result []*Category) {

	return
}
