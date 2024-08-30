package user

import (
	"GinDemo/middleware/Accessable"
	"GinDemo/middleware/jwt"
	"github.com/gin-gonic/gin"
)

func init() {

}

func Init(gin *gin.Engine) {
	group := gin.Group("user")
	group.Use(jwt.SimpleJwtAuthMiddleware(), Accessable.AccessableMiddleware())
	{
		group.POST("", Add)
		group.DELETE("/:id", Delete)
		group.PUT("/:id", Update)
		group.GET("", Get)
		group.GET("/:id", ById)
	}
}
