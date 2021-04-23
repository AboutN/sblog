package config

import (
	"bytes"
	_ "embed"

	"sblog/log"

	"github.com/spf13/viper"
)

var (
	//go:embed app.conf
	defconf []byte

	BlogName string // 博客名称
	Port     string // sblog运行端口
	Theme    string // 主题名称
	PageSize int    // 分页大小
	DBPath   string // 数据库文件路径
)

const ConfigFile = "app.conf"

// 则使用默认配置信息
func Init() {
	viper.SetConfigType("yaml")
	err := viper.ReadConfig(bytes.NewBuffer(defconf))
	if err != nil {
		log.Error().Msg(err.Error())
		panic(err)
	}

	BlogName = viper.GetString("BlogName")
	Port = viper.GetString("Port")
	Theme = viper.GetString("Theme")
	PageSize = viper.GetInt("PageSize")
	DBPath = viper.GetString("DBPath")
}
