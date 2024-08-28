package user

import (
	"GinDemo/db"
	"GinDemo/uilty"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 增加用户
func Add(c *gin.Context) {
	var user User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusOK, uilty.ErrorResponseDefault{Message: "缺少参数", Code: -1})
		return
	}
	err := db.Create(&user)
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
	user := User{Id: &num}
	_, er := db.FindById(user, num)
	if er != nil {
		c.JSON(http.StatusOK, uilty.ErrorResponseDefault{Message: "用户不存在", Code: -1})
		return
	}
	er = db.DeleteById(user, user.IdValue())
	if er != nil {
		c.JSON(http.StatusOK, uilty.ErrorResponseDefault{Message: "删除用户失败", Code: -1})
		return
	}
	c.JSON(http.StatusOK, uilty.SuccessResponse[string](nil))
}

// 跟新用户
func Update(c *gin.Context) {
	var reQuestUser User
	if err := c.ShouldBind(&reQuestUser); err != nil {
		c.JSON(http.StatusOK, uilty.ErrorResponseDefault{Message: "缺少参数", Code: -1})
		return
	}
	if !reQuestUser.isValidWhenUpdate() {
		c.JSON(http.StatusOK, uilty.ErrorResponseDefault{Message: "缺少参数", Code: -1})
		return
	}
	id := c.Param("id")
	num, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusOK, uilty.ErrorResponseDefault{Message: err.Error(), Code: -1})
		return
	}
	user, er := db.FindById(reQuestUser, num)
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
	er = db.UpdateById(user)
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
		pageCxt := db.PageContext{Page: p, PageSize: m}
		retval, err := db.PageFind[User](pageCxt)
		if err != nil {
			c.JSON(http.StatusOK, uilty.ErrorResponseDefault{Message: err.Error(), Code: -1})
			return
		}
		count, er := db.Total(&User{})
		if er != nil {
			c.JSON(http.StatusOK, uilty.ErrorResponseDefault{Message: err.Error(), Code: -1})
		}
		ret := struct {
			Total any    `json:"total"`
			Users []User `json:"users"`
		}{
			Total: count,
			Users: retval,
		}

		c.JSON(http.StatusOK, uilty.SuccessResponse(&ret))
		return
	}
	if maxNum == "" {
		result, err1 := db.FindByLimit[User](-1)
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
	result, err1 := db.FindByLimit[User](num)
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
	find := User{Id: &num}
	user, er := db.FindById(find, find.IdValue())
	if er != nil {
		c.JSON(http.StatusOK, uilty.ErrorResponseDefault{Message: "用户不存在", Code: -1})
		return
	}
	c.JSON(http.StatusOK, uilty.SuccessResponse(&user))
}
