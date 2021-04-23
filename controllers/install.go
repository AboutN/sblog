package controllers

import (
	"fmt"
	"net/http"
	"sblog/log"

	"github.com/gin-gonic/gin"
	// "github.com/gin-gonic/gin/binding"
	// "github.com/go-playground/locales/en"
	// "github.com/go-playground/locales/zh"
	// ut "github.com/go-playground/universal-translator"
	// "github.com/go-playground/validator/v10"
	// enTranslations "github.com/go-playground/validator/v10/translations/en"
	// chTranslations "github.com/go-playground/validator/v10/translations/zh"
)

// var trans ut.Translator

// func transInit(local string) (err error) {
// 	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
// 		zhT := zh.New() //chinese
// 		enT := en.New() //english
// 		uni := ut.New(enT, zhT, enT)

// 		var o bool
// 		trans, o = uni.GetTranslator(local)
// 		if !o {
// 			return fmt.Errorf("uni.GetTranslator(%s) failed", local)
// 		}
// 		//register translate
// 		// 注册翻译器
// 		switch local {
// 		case "en":
// 			err = enTranslations.RegisterDefaultTranslations(v, trans)
// 		case "zh":
// 			err = chTranslations.RegisterDefaultTranslations(v, trans)
// 		default:
// 			err = enTranslations.RegisterDefaultTranslations(v, trans)
// 		}
// 		return
// 	}
// 	return
// }

type InitConfig struct {
	BlogName        string `form:"blog-name" binding:"required"`                         // 博客名称
	Theme           string `form:"theme" binding:"required"`                             // 主题
	Port            int    `form:"port" binding:"required"`                              // 端口
	PageSize        int    `form:"page-size" binding:"required"`                         // 分页
	Email           string `form:"email" binding:"required"`                             // 邮箱
	NickName        string `form:"nick-name" binding:"required"`                         // 昵称
	Password        string `form:"password" binding:"required"`                          // 密码
	ConfirmPassword string `form:"confirm-password" binding:"required,eqfield=Password"` // 确认密码
}

// Step1 初始化信息页面
func Step1(ctx *gin.Context) {
	fmt.Println("STEP1")
	ctx.HTML(http.StatusOK, "install/step1.htm", nil)

}

func Step1Post(ctx *gin.Context) {
	// transInit("zh")
	data := InitConfig{}
	err := ctx.ShouldBind(&data)
	if err != nil {
		// valerr, ok := err.(validator.ValidationErrors)
		// if !ok {
		// 	// 非validator.ValidationErrors类型错误直接返回
		// 	ctx.JSON(http.StatusOK, gin.H{
		// 		"msg": err.Error(),
		// 	})
		// 	return
		// }
		// // validator.ValidationErrors类型错误则进行翻译
		// ctx.JSON(http.StatusOK, gin.H{
		// 	"msg": valerr.Translate(trans),
		// })
		log.Info().Msg(err.Error())
		fmt.Println(err.Error())
		ctx.HTML(http.StatusOK, "install/step1.htm", nil)
		return
	}

	ctx.Redirect(http.StatusOK, "/finish")
}

// Step2 结果页面
func Step2(ctx *gin.Context) {

	// 创建配置文件
	ctx.HTML(http.StatusOK, "install/step2.htm", nil)
	// 创建数据库并初始化
}
