package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Step1 初始化信息页面
func Step1(ctx *gin.Context) {
	// 判断配置文件是否存在, 如果配置文件已经存在则无法进入该页面
	// 配置文件只能保存在和应用程序同级目录, 名字为 app.conf
	ctx.HTML(http.StatusOK, "step1.htm", nil)
}

// Step2 结果页面
func Step2(ctx *gin.Context) {
	// 创建配置文件

	// 创建数据库并初始化
}
