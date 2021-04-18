package controllers

import (
	"net/http"
	"sblog/entities"
	"sblog/scode"
	"sblog/services"

	"github.com/gin-gonic/gin"
)

// Login 游客登录
func Login(ctx *gin.Context) {

}

// Manage  博客管理
func Manage(ctx *gin.Context) {
	entity := &entities.ManageLogin{}

	// 数据绑定及验证
	err := ctx.ShouldBindJSON(entity)
	if err != nil {
		//
		ctx.JSON(http.StatusOK, entities.ResultEntity{
			Code: scode.Failure,
			Info: "数据错误",
		})
		return
	}

	if entity.UserID == "" || entity.Password == "" {
		ctx.JSON(http.StatusOK, entities.ResultEntity{
			Code: scode.Failure,
			Info: "数据错误",
		})
		return
	}

	result := services.ManageService(entity.UserID, entity.Password)

	ctx.JSON(http.StatusOK, result)
}
