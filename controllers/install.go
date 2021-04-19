package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type InitConfig struct {
	BlogName        string `form:"blog-name"`        // 博客名称
	Theme           string `form:"theme"`            // 主题
	Port            int    `form:"port"`             // 端口
	PageSize        int    `form:"page-size"`        // 分页
	Email           string `form:"email"`            // 邮箱
	NickName        string `form:"nick-name"`        // 昵称
	Password        string `form:"password"`         // 密码
	ConfirmPassword string `form:"confirm-password"` // 确认密码
}

// Step1 初始化信息页面
func Step1(ctx *gin.Context) {
	fmt.Println("STEP1")
	ctx.HTML(http.StatusOK, "install/step1.htm", nil)

}

func Step1Post(ctx *gin.Context) {

	ctx.HTML(http.StatusOK, "install/step1.htm", nil)
}

// Step2 结果页面
func Step2(ctx *gin.Context) {
	fmt.Println("STEP2")
	// 创建配置文件
	ctx.HTML(http.StatusOK, "install/step2.htm", nil)
	// 创建数据库并初始化
}
