package main

import (
	"GinDemo/db/mysql"
	signin "GinDemo/signIn"
	"GinDemo/user"

	"github.com/gin-gonic/gin"
)

func main() {
	gin := gin.Default()
	mysql.Init(gin)
	user.Init(gin)
	signin.Init(gin)
	gin.Run("127.0.0.1:8080")
}
