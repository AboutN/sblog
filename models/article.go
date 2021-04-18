package models

import (
	"sblog/log"
	"time"
)

// Article 文章
type Article struct {
	ID         string    // 编号
	Title      string    // 标题
	Content    string    // 内容
	Overview   string    // 概述
	CreateTime time.Time // 创建时间
	UpdateTime time.Time // 更新时间
	IsOriginal bool      // 是否原创
	IsTop      bool      // 是否置顶
	IsPrivacy  bool      // 仅自己可见
}

// Create 添加新的文章
func (m *Article) Create() (result bool) {
	err := conn.Create(m).Error
	if err != nil {
		result = false
		log.Error().Err(err)
		return
	}
	result = true
	return
}

// Get 根据文章ID 获取文章
func (m *Article) Get() (result *Article) {
	db := conn.First(result)
	err := db.Error
	if err != nil {
		result = nil
		log.Error().Err(err)
		return
	}
	if db.RowsAffected == 0 {
		result = nil
		log.Warn().Msg("获取文章失败")
		return
	}
	return
}

// Edit 编辑文章
func (m *Article) Edit(param Params) (result bool) {
	return
}
