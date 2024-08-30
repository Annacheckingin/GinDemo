package user

import (
	"GinDemo/db/mysql"
	"GinDemo/model"
	"GinDemo/uilty"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 增加用户
func Add(c *gin.Context) {
	var user model.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusOK, uilty.ErrorResponseDefault{Message: "缺少参数", Code: -1})
		return
	}
	err := mysql.Create(&user)
	if err != nil {
		c.JSON(http.StatusOK, uilty.ErrorResponseDefault{Message: "添加用户失败" + err.Error(), Code: -1})
		return
	}
	c.JSON(http.StatusOK, uilty.SuccessResponse[string](nil))
}

// 删除用户
func Delete(c *gin.Context) {
	id := c.Param("id")
	num, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusOK, uilty.ErrorResponseDefault{Message: err.Error(), Code: -1})
		return
	}
	user := model.User{Id: &num}
	_, er := mysql.FindById(user, num)
	if er != nil {
		c.JSON(http.StatusOK, uilty.ErrorResponseDefault{Message: "用户不存在", Code: -1})
		return
	}
	er = mysql.DeleteById(user, user.IdValue())
	if er != nil {
		c.JSON(http.StatusOK, uilty.ErrorResponseDefault{Message: "删除用户失败", Code: -1})
		return
	}
	c.JSON(http.StatusOK, uilty.SuccessResponse[string](nil))
}

// 跟新用户
func Update(c *gin.Context) {
	var reQuestUser model.User
	if err := c.ShouldBind(&reQuestUser); err != nil {
		c.JSON(http.StatusOK, uilty.ErrorResponseDefault{Message: "缺少参数", Code: -1})
		return
	}
	if !reQuestUser.IsValidWhenUpdate() {
		c.JSON(http.StatusOK, uilty.ErrorResponseDefault{Message: "缺少参数", Code: -1})
		return
	}
	id := c.Param("id")
	num, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusOK, uilty.ErrorResponseDefault{Message: err.Error(), Code: -1})
		return
	}
	user, er := mysql.FindById(reQuestUser, num)
	if er != nil {
		c.JSON(http.StatusOK, uilty.ErrorResponseDefault{Message: "用户不存在", Code: -1})
		return
	}
	if reQuestUser.Name != nil {
		user.Name = reQuestUser.Name
	}
	if reQuestUser.Password != nil {
		user.Password = reQuestUser.Password
	}
	er = mysql.UpdateById(user)
	if er != nil {
		c.JSON(http.StatusOK, uilty.ErrorResponseDefault{Message: "更新失败" + er.Error(), Code: -1})
		return
	}
	c.JSON(http.StatusOK, uilty.SuccessResponse[string](nil))
}

// Get 获取前N条数据
func Get(c *gin.Context) {
	maxNum := c.Query("count")
	page := c.Query("page")
	if len(page) != 0 && len(maxNum) != 0 {
		m, err := strconv.Atoi(maxNum)
		p, err := strconv.Atoi(page)
		if err != nil {
			c.JSON(http.StatusOK, uilty.ErrorResponseDefault{Message: err.Error(), Code: -1})
			return
		}
		pageCxt := mysql.PageContext{Page: p, PageSize: m}
		retval, err := mysql.PageFind[model.User](pageCxt)
		if err != nil {
			c.JSON(http.StatusOK, uilty.ErrorResponseDefault{Message: err.Error(), Code: -1})
			return
		}
		count, er := mysql.Total(&model.User{})
		if er != nil {
			c.JSON(http.StatusOK, uilty.ErrorResponseDefault{Message: err.Error(), Code: -1})
			return
		}
		ret := struct {
			Total any          `json:"total"`
			Users []model.User `json:"users"`
		}{
			Total: count,
			Users: retval,
		}

		c.JSON(http.StatusOK, uilty.SuccessResponse(&ret))
		return
	}
	if maxNum == "" {
		result, err1 := mysql.FindByLimit[model.User](-1)
		if err1 != nil {
			c.JSON(http.StatusOK, uilty.ErrorResponseDefault{Message: err1.Error()})
			return
		}
		c.JSON(http.StatusOK, uilty.SuccessResponseArray(&result))
		return
	}
	num, err := strconv.Atoi(maxNum)
	if err != nil {
		c.JSON(http.StatusOK, uilty.ErrorResponseDefault{Message: err.Error(), Code: -1})
		return
	}
	result, err1 := mysql.FindByLimit[model.User](num)
	if err1 != nil {
		c.JSON(http.StatusOK, uilty.ErrorResponseDefault{Message: err1.Error()})
		return
	}
	c.JSON(http.StatusOK, uilty.SuccessResponseArray(&result))
}

// 获取单个用户
func ById(c *gin.Context) {
	id := c.Param("id")
	num, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusOK, uilty.ErrorResponseDefault{Message: err.Error(), Code: -1})
		return
	}
	find := model.User{Id: &num}
	user, er := mysql.FindById(find, find.IdValue())
	if er != nil {
		c.JSON(http.StatusOK, uilty.ErrorResponseDefault{Message: "用户不存在", Code: -1})
		return
	}
	c.JSON(http.StatusOK, uilty.SuccessResponse(&user))
}
