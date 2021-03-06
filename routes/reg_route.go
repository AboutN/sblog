package routes

import (
	"sblog/controllers"
	"sblog/routes/middleware"

	"github.com/gin-gonic/gin"
)

func RegRout(router *gin.Engine) {

	installd := router.Group("", middleware.IsInstalld)
	{
		// 系统初始化
		installd.GET("/install", controllers.Step1)
		installd.POST("/install", controllers.Step1Post)
		// 初始化完成
		installd.GET("/finish", controllers.Step2)
	}

	// 注册
	// router.GET("/registry")
	// 登录
	router.GET("/login", controllers.Login)
	// 登出
	// router.GET("/logout", controllers.Logout)
	// 管理
	router.GET("/manage", controllers.Manage)

	// RSS
	router.GET("/rss", nil)

	// 加载更多文章
	router.GET("/next/:n", nil)

	// 未知路径
	router.NoRoute()

	manage := router.Group("")
	{
		// 首页
		manage.GET("/", nil)
		// 分类列表
		manage.GET("/tag", nil)
		// 添加分类
		manage.GET("/tag/add", nil)
		// 个人信息
		manage.GET("/personal", nil)

		// 写博客
		manage.GET("/write/*id", nil)
		// 博文
		manage.GET("/p/:id", nil)
		// 删除
		manage.GET("/del/:id", nil)
	}

}
