package templates

import (
	"fmt"
	"sblog/config"

	"github.com/gin-gonic/gin"
)

// LoadTemplates 加载模板
func LoadTemplates(route *gin.Engine) {
	dir := fmt.Sprintf("themes/%s/layout/*", config.Theme)
	route.LoadHTMLGlob(dir)
}
