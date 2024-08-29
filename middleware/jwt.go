package middleware

import (
	"GinDemo/uilty"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type User struct {
	Name     *string `json:"name"`
	Password *string `json:"password"`
}

func Init(gin *gin.Engine) {

}

func MakeJWT() *jwt.GinJWTMiddleware {
	return initParams()
}

func Validate(u *User) bool {
	return false
}

func initParams() *jwt.GinJWTMiddleware {

	return &jwt.GinJWTMiddleware{
		Realm:           "test zone",
		Key:             []byte("secret key"),
		Timeout:         time.Hour,
		MaxRefresh:      time.Hour,
		IdentityKey:     "name",
		IdentityHandler: identityHandler(),
		PayloadFunc:     payloadFunc(),
		Authenticator:   authenticator(),
		Unauthorized:    unauthorized(),
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
	}
}

func identityHandler() func(*gin.Context) interface{} {
	return func(c *gin.Context) interface{} {
		claims := jwt.ExtractClaims(c)
		name := claims["name"].(string)
		password := claims["password"].(string)
		ret := struct {
			Name     *string `json:"name"`
			Password *string `json:"password"`
		}{
			Name:     &name,
			Password: &password,
		}
		return &ret
	}
}

func payloadFunc() func(data interface{}) jwt.MapClaims {
	return func(data interface{}) jwt.MapClaims {
		if v, ok := data.(*User); ok {
			return jwt.MapClaims{
				"name": v.Name,
			}
		}
		return jwt.MapClaims{}
	}
}

func authenticator() func(c *gin.Context) (interface{}, error) {
	return func(c *gin.Context) (interface{}, error) {

		var loginVals User
		if err := c.ShouldBind(&loginVals); err != nil {
			return "", jwt.ErrMissingLoginValues
		}
		name := loginVals.Name
		password := loginVals.Password
		if name == nil || password == nil {
			return nil, jwt.ErrMissingLoginValues
		}
		if len(*name) == 0 || len(*password) == 0 {
			return nil, jwt.ErrFailedAuthentication
		}
		return User{
			Name:     name,
			Password: password,
		}, nil
	}
}

func unauthorized() func(c *gin.Context, code int, message string) {
	return func(c *gin.Context, code int, message string) {
		c.JSON(code, uilty.ErrorResponseDefault{Message: "token校验失败", Code: -1})
	}
}
