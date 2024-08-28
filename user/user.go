package user

import "github.com/gin-gonic/gin"

func init() {

}

func Init(gin *gin.Engine) {
	group := gin.Group("user")
	{
		group.POST("", Add)
		group.DELETE("/:id", Delete)
		group.PUT("/:id", Update)
		group.GET("", Get)
		group.GET("/:id", ById)
	}
}
