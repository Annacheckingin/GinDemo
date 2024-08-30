package Accessable

import (
	"GinDemo/db/noSql"
	"GinDemo/middleware/jwt"
	"GinDemo/uilty"
	"github.com/gin-gonic/gin"
	"github.com/jefferyjob/go-easy-utils/v2/anyUtil"
	_ "github.com/jefferyjob/go-easy-utils/v2/sliceUtil"
)

func AccessableMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtSub, ok := c.Get(jwt.JWT_SUB_CONTEXT_KEY)
		if !ok {
			uilty.ErrorMessage(c, "没有权限")
			c.Abort()
			return
		}
		sub := anyUtil.AnyToStr(jwtSub)
		if !ok {
			uilty.ErrorMessage(c, "token格式有误")
			c.Abort()
			return
		}
		value := noSql.Get(sub)
		if value == nil || len(*value) == 0 {
			uilty.ErrorMessage(c, "token已过期")
			c.Abort()
			return
		}
		c.Next()
	}

}
