package signin

import (
	"GinDemo/middleware/jwt"
	"github.com/gin-gonic/gin"
)

func Init(gin *gin.Engine) {
	group := gin.Group("sign")
	group.PUT("", SignIn)
	group.POST("", SignUp)
	group.DELETE("", Logout, jwt.SimpleJwtAuthMiddleware())
	group.PATCH("", Quit, jwt.SimpleJwtAuthMiddleware())
}
