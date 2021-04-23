package models

import (
	"sblog/log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var conn *gorm.DB

func Init() {
	db, err := gorm.Open(sqlite.Open("sblog.db"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			// TablePrefix:   "t_",                              // 表名前缀，`User`表为`t_users`
			SingularTable: true, // 使用单数表名，启用该选项后，`User` 表将是`user`
		},
	})
	if err != nil {
		log.Error().Msg(err.Error())
		panic("err: 数据库连接失败")
	}
	db.AutoMigrate(&Profile{})
	conn = db
}

// Params 参数列表
type Params map[string]interface{}
