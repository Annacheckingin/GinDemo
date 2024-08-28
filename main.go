package main

import (
	"GinDemo/db"
	"GinDemo/user"

	"github.com/gin-gonic/gin"
)

func main() {
	gin := gin.Default()
	db.Init(gin)
	user.Init(gin)
	gin.Run("127.0.0.1:8080")
}
