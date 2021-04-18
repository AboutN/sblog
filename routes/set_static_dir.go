package routes

import (
	"fmt"

	"sblog/config"

	"github.com/gin-gonic/gin"
)

func SetStaticDir(route *gin.Engine) {
	dir := fmt.Sprintf("themes/%s/assets", config.Theme)
	route.Static("/theme", dir)
}
