package model

import (
	"github.com/jinzhu/gorm"
)

// UserRole 用户-角色关联表
type UserRole struct {
	gorm.Model
	UserID uint `gorm:"not null;size:16" json:"user_id"`
	RoleID uint `gorm:"not null;size:16" json:"role_id"`
}
