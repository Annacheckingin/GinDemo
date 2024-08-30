package signin

import (
	"GinDemo/middleware/Accessable"
	"GinDemo/middleware/jwt"
	"github.com/gin-gonic/gin"
)

func Init(gin *gin.Engine) {
	group := gin.Group("sign")
	//登录
	group.PUT("", SignIn)
	//注册
	group.POST("", SignUp)
	//注销
	group.DELETE("", Logout, jwt.SimpleJwtAuthMiddleware(), Accessable.AccessableMiddleware())
	//退出登录
	group.PATCH("", Quit, jwt.SimpleJwtAuthMiddleware(), Accessable.AccessableMiddleware())
}
