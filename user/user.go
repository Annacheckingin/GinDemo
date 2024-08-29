package user

import (
	"GinDemo/middleware"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func init() {

}

func Init(gin *gin.Engine) {
	group := gin.Group("user")
	jwtAuth := middleware.MakeJWT()
	group.Use(handlerMiddleWare(jwtAuth))
	{
		group.POST("", Add)
		group.DELETE("/:id", Delete)
		group.PUT("/:id", Update)
		group.GET("", Get)
		group.GET("/:id", ById)
	}
}

func handlerMiddleWare(authMiddleware *jwt.GinJWTMiddleware) gin.HandlerFunc {
	return func(context *gin.Context) {
		errInit := authMiddleware.MiddlewareInit()
		if errInit != nil {
			panic("JWt middleware init error")
		}
	}
}
