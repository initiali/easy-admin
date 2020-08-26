package model

import (
	"github.com/jinzhu/gorm"
)

// Role 系统角色表
type Role struct {
	gorm.Model
	RoleName string `gorm:"size:32;not null"`
	RoleFlag string `gorm:"size:32;not null"`
	Sort     int    `gorm:"size:4;not null"`
}
