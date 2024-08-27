package user

import (
	_ "gorm.io/gorm"
	"time"
)

type User struct {
	Name      *string   `gorm:"column:user_name" json:"name"`
	Id        int       `gorm:"column:user_id;primary;autoincrement" json:"-"`
	Password  *string   `gorm:"column:password" json:"password" binding:"required"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"-"` // 创建时间，由 GORM 自动管理
}
