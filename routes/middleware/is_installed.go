package middleware

import (
	"fmt"
	"net/http"
	"sblog/config"
	"sblog/utils"
	"time"

	"github.com/gin-gonic/gin"
)

// IsInstalld 检测是否已经执行过安装程序
// 通过判断app.conf文件是否存在来确定是否执行过安装程序
func IsInstalld(ctx *gin.Context) {
	fmt.Println(time.Now())
	if utils.IsExist(config.ConfigFile) {
		ctx.HTML(http.StatusOK, "error.htm", nil)
		ctx.Abort()
	}
}
