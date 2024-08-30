package signin

import (
	"GinDemo/db/mysql"
	"GinDemo/db/noSql"
	_ "GinDemo/db/noSql"
	"GinDemo/middleware/jwt"
	"GinDemo/model"
	"GinDemo/uilty"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"time"
)

const expireSeconds time.Duration = 60 * 60 * time.Second

func SignIn(c *gin.Context) {
	usr := model.User{}
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
		User          model.User `json:"user"`
		Authorization string     `json:"Authorization"`
	}{
		User:          usr,
		Authorization: token,
	}
	uilty.DoneWithReturn(c, ret)
}

func SignUp(c *gin.Context) {
	usr := model.User{}
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

func insertNewUserRecord(gin *gin.Context, user model.User) (*model.User, error) {
	var count int64
	result := mysql.Db.Where("user_name = ?", user.Name).Find(&model.User{}).Count(&count)
	if result.Error != nil {
		return nil, result.Error
	}
	if count != 0 {
		return nil, fmt.Errorf("用户已存在")
	}
	er := mysql.Db.Create(&user)
	if er != nil {
		return nil, er.Error

	}
	return &user, nil
}

func signIn(user model.User) (string, error) {
	token, er := jwt.SimpleJwt(expireSeconds, *user.Name)
	if er != nil {
		return "", er
	}
	if len(*user.Name) == 0 {
		return "", fmt.Errorf("用户名为空")
	}
	er = noSql.SetString(*user.Name, token, expireSeconds)
	if er != nil {
		return "", er
	}
	return token, nil
}

func Quit(c *gin.Context) {
	usr := model.User{}
	if er := c.ShouldBindBodyWith(&usr, binding.JSON); er != nil {
		uilty.ErrorMessage(c, "传参有误")
		return
	}
	if (len(*usr.Name)) == 0 {
		uilty.ErrorMessage(c, "传参有误")
		return
	}
	er := noSql.RemoveString(*usr.Name)
	if er != nil {
		uilty.Error(c, er)
	}
	uilty.Done(c)
}

func Logout(c *gin.Context) {
	usr := model.User{}
	if er := mysql.DeleteById(usr, usr.IdValue()); er != nil {
		uilty.Error(c, er)
		return
	}
	Quit(c)
}
