package model

import (
	"github.com/jinzhu/gorm"
)

// Menu 菜单表
type Menu struct {
	gorm.Model
	MenuPid   uint   `gorm:"size:16;not null"`
	MenuPids  string `gorm:"size:16;not null"`
	IsLeaf    int    `gorm:"size:4;not null"`
	Name      string `gorm:"size:32;not null"`
	URL       string `gorm:"size:32;not null"`
	Icon      string `gorm:"size:32;not null"`
	IconColor string `gorm:"size:16;not null"`
	Sort      int    `gorm:"size:4;not null"`
	Level     int    `gorm:"size:4;not null"`
	Status    int    `gorm:"size:4;not null"`
}
