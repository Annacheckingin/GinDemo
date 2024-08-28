package user

import (
	"time"

	_ "gorm.io/gorm"
)

type User struct {
	Name      *string   `gorm:"column:user_name" json:"name"`
	Id        *int      `gorm:"column:user_id;primaryKey;autoincrement" json:"id"`
	Password  *string   `gorm:"column:password" json:"password"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"-"` // 创建时间，由 GORM 自动管理
}

func (u User) IdValue() any {
	return u.Id
}

func (User) TableName() string {
	return "users"
}

func (u User) isValidWhenUpdate() bool {
	if u.Name == nil && u.Password == nil {
		return false
	}
	return true
}
