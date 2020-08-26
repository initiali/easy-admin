package model

import (
	"github.com/jinzhu/gorm"
)

// RoleMenu 角色-菜单关联表
type RoleMenu struct {
	gorm.Model
	RoleID uint `gorm:"size:16;not null"`
	MenuID uint `gorm:"size:16;not null"`
}
