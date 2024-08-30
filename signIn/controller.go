package signin

import (
	"GinDemo/db/mysql"
	"GinDemo/db/noSql"
	_ "GinDemo/db/noSql"
	"GinDemo/middleware/jwt"
	"GinDemo/uilty"
	"GinDemo/user"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"time"
)

func SignIn(c *gin.Context) {
	usr := user.User{}
	if er := c.ShouldBindBodyWith(&usr, binding.JSON); er != nil {
		uilty.ErrorMessage(c, er.Error())
		return
	}
	token, er := signIn(usr)
	if er != nil {
		uilty.Error(c, er)
		return
	}
	ret := struct {
		User          user.User `json:"user"`
		Authorization string    `json:"Authorization"`
	}{
		User:          usr,
		Authorization: token,
	}
	uilty.DoneWithReturn(c, ret)
}

func SignUp(c *gin.Context) {
	usr := user.User{}
	if er := c.ShouldBindBodyWith(&usr, binding.JSON); er != nil {
		uilty.ErrorMessage(c, "传参有误")
		return
	}
	_, er := insertNewUserRecord(c, usr)
	if er != nil {
		uilty.Error(c, er)
		return
	}
	SignIn(c)
}

func insertNewUserRecord(gin *gin.Context, user user.User) (user.User, error) {
	er := mysql.Create(&user)
	return user, er
}

func signIn(user user.User) (string, error) {
	token, er := jwt.SimpleJwt(60 * time.Second)
	if er != nil {
		return "", er
	}
	if len(*user.Name) == 0 {
		return "", fmt.Errorf("用户名为空")
	}
	er = noSql.SetString(*user.Name, token, 60*time.Second)
	if er != nil {
		return "", er
	}
	return token, nil
}

func Quit(c *gin.Context) {

}

func Logout(c *gin.Context) {

}
