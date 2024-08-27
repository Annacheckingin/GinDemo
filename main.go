package main

import (
	"GinDemo/db"
	"GinDemo/user"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	db.Setup()
	group := router.Group("user")
	{
		group.POST("", user.Add)
		group.DELETE("/:id", user.Delete)
		group.PUT("/:id", user.Update)
		group.GET("", user.Get)
		group.GET("/:id", user.ById)
	}
	router.Run("127.0.0.1:8080")

}
