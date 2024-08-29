package middleware

import (
	"GinDemo/user"
	"reflect"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func Init(gin *gin.Engine) {

}

func JWT(u *user.User) *jwt.GinJWTMiddleware {
	return initParams(u)
}

func Validate(u *user.User) bool {
	return false
}

func initParams(u *user.User) *jwt.GinJWTMiddleware {

	return &jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: "id",
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	}
}

func identityHandler(u *user.User) func(*gin.Context) (interface{}, error) {
	return func(c *gin.Context) (interface{}, error) {
		t := reflect.TypeOf(u)
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			jsonTag := field.Tag.Get("json")
			if jsonTag == "name" {
				value := reflect.ValueOf(u).Field(i).Interface()
				return value, nil
			}
			return nil, nil
		}
		return nil, nil
	}
}
