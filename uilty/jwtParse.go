package uilty

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func JwtTokenFecth(c *gin.Context) *string {
	payload := c.GetHeader("Authorization")
	parts := strings.Split(payload, " ")
	if len(parts) != 2 {
		return nil
	}
	return &parts[1]
}
