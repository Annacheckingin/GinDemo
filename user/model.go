package user

import (
	"time"

	_ "gorm.io/gorm"
)

type User struct {
	Name      *string   `gorm:"column:user_name" json:"name"`
	Id        *int      `gorm:"column:user_id;primary;autoincrement" json:"id"`
	Password  *string   `gorm:"column:password" json:"password" binding:"required"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"-"` // 创建时间，由 GORM 自动管理
}

func (User) TableName() string {
	return "users"
}
