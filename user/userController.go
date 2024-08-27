package user

import (
	"GinDemo/db"
	"GinDemo/uilty"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Add(c *gin.Context) {
	var user User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusOK, uilty.ErrorResponseDefault{Message: "缺少参数", Code: -1})
		return
	}
	result := db.Db.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusOK, uilty.ErrorResponseDefault{Message: "添加用户失败" + result.Error.Error(), Code: -1})
		return
	}
	c.JSON(http.StatusOK, uilty.SuccessResponse[string](nil))
}

func Delete(c *gin.Context) {
	id := c.Param("id")
	queryRt := db.Db.Where("user_id = ?", id).First(&User{})
	if queryRt.Error != nil {
		c.JSON(http.StatusOK, uilty.ErrorResponseDefault{Message: "用户不存在", Code: -1})
		return
	}
	deleteRt := db.Db.Where("user_id = ?", id).Delete(&User{})
	if deleteRt.Error != nil {
		c.JSON(http.StatusOK, uilty.ErrorResponseDefault{Message: "删除用户失败", Code: -1})
		return
	}
	c.JSON(http.StatusOK, uilty.SuccessResponse[string](nil))
}

func Update(c *gin.Context) {
	id := c.Param("id")
	num, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusOK, uilty.ErrorResponseDefault{Message: err.Error(), Code: -1})
		return
	}

	if dbById(num) == nil {
		c.JSON(http.StatusOK, uilty.ErrorResponseDefault{Message: "用户不存在", Code: -1})
		return
	}

	var user User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusOK, uilty.ErrorResponseDefault{Message: "缺少参数", Code: -1})
		return
	}

	if user.Name != nil {
		if len(*user.Name) > 0 {
			result := db.Db.Model(&User{}).Where("user_id = ?", num).Update("user_name", user.Name)
			if result.Error != nil {
				c.JSON(http.StatusOK, uilty.ErrorResponseDefault{Message: "更新失败" + result.Error.Error(), Code: -1})
				return
			}
			c.JSON(http.StatusOK, uilty.SuccessResponse[string](nil))
		}
	}

	if user.Password != nil {
		if len(*user.Password) > 0 {
			result := db.Db.Model(&User{}).Where("user_id = ?", num).Update("password", user.Password)
			if result.Error != nil {
				c.JSON(http.StatusOK, uilty.ErrorResponseDefault{Message: "更新失败" + result.Error.Error(), Code: -1})
				return
			}
			c.JSON(http.StatusOK, uilty.SuccessResponse[string](nil))
		}
	}

}

func Get(c *gin.Context) {
	var users []User
	result := db.Db.Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusOK, uilty.ErrorResponseDefault{Message: result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, uilty.SuccessResponseArray(&users))
}

func ById(c *gin.Context) {
	id := c.Param("id")
	num, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusOK, uilty.ErrorResponseDefault{Message: err.Error(), Code: -1})
		return
	}
	user := dbById(num)
	if user == nil {
		c.JSON(http.StatusOK, uilty.ErrorResponseDefault{Message: "用户不存在", Code: -1})
		return
	}
	c.JSON(http.StatusOK, uilty.SuccessResponse(dbById(num)))
}

func dbById(id int) *User {
	var user User
	if err := db.Db.Where("user_id = ?", id).First(&user).Error; err != nil {
		return nil
	}
	return &user
}
