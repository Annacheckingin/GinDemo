package middleware

import (
	"GinDemo/uilty"
	"fmt"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type JWTExample struct {
	Sub  string `json:"sub"`
	Name string `json:"name"`
}

func Init(gin *gin.Engine) {

}

func MakeJWT() *jwt.GinJWTMiddleware {
	return initParams()
}

func initParams() *jwt.GinJWTMiddleware {

	return &jwt.GinJWTMiddleware{
		Realm:            "test zone",
		SigningAlgorithm: "HS256",
		Key:              []byte("secret key"),
		Timeout:          time.Hour,
		MaxRefresh:       time.Hour,
		IdentityHandler:  identityHandler(),
		PayloadFunc:      payloadFunc(),
		Authorizator:     authorizator(),
		Authenticator:    authenticator(),
		Unauthorized:     unauthorized(),
		TokenHeadName:    "Bearer",
		TimeFunc:         time.Now,
	}
}

func identityHandler() func(*gin.Context) interface{} {
	return func(c *gin.Context) interface{} {
		claims := jwt.ExtractClaims(c)
		name := claims["name"].(string)
		sub := claims["sub"].(string)
		ret := JWTExample{Sub: sub, Name: name}
		return &ret
	}
}

// 对jwt进行校验——时间和完整性等基本校验已经提前完成
func authorizator() func(data interface{}, c *gin.Context) bool {
	return func(data interface{}, c *gin.Context) bool {
		if _, ok := data.(*JWTExample); ok {
			return true
		}
		return true
	}
}

func payloadFunc() func(data interface{}) jwt.MapClaims {
	return func(data interface{}) jwt.MapClaims {
		if v, ok := data.(*JWTExample); ok {
			/**
			   "sub": "1234567890",
			  "name": "John Doe",
			  "exp":1735689600
			*/
			return jwt.MapClaims{
				"name": v.Name,
				"sub":  v.Sub,
			}
		}
		return jwt.MapClaims{}
	}
}

// MARK: 只有使用了LoginHandler才触发此方法，此方法的返回值用于jwt生成，放置于jwt的负载区域；这部分数据的索引设置为GinJWTMiddleware的
// identityKey所设置的字符串——即为jwt的payload当中同·sub·这样的key同级的存在
func authenticator() func(c *gin.Context) (interface{}, error) {
	return func(c *gin.Context) (interface{}, error) {
		return nil, nil
	}
}

func unauthorized() func(c *gin.Context, code int, message string) {
	return func(c *gin.Context, code int, message string) {
		fmt.Printf("Code is %d. message is %s", code, message)
		c.JSON(http.StatusOK, uilty.ErrorResponseDefault{Message: "token校验失败", Code: -1})
	}
}
