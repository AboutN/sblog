package main

import (
	"sblog/config"
	"sblog/log"
	"sblog/models"
	"sblog/routes"
	"sblog/templates"

	"github.com/gin-gonic/gin"
)

var blogsvr *WebServer

func main() {
	// 启动web服务
	blogsvr.Run()
}

// 初始化应用
func init() {

	// 初始化日志工具
	log.Init()

	// 加载配置信息
	config.Init()

	// 创建数据库连接
	models.Init()

	// 初始化web服务
	blogsvr = &WebServer{}
	blogsvr.Init()
}

type WebServer struct {
	engine *gin.Engine
}

func (ws *WebServer) Run() {
	ws.engine.Run(":" + config.Port)

}

func (ws *WebServer) Init() {
	ws.engine = gin.Default()
	// 加载模板
	templates.LoadTemplates(ws.engine)
	// 设置静态目录
	routes.SetStaticDir(ws.engine)
	// 注册路由
	routes.RegRout(ws.engine)
}
