package model

import (
	"github.com/jinzhu/gorm"
)

// Org 系统组织结构表
type Org struct {
	gorm.Model
	OrgPid  uint   `gorm:"size:16;not null"`
	OrgPids string `gorm:"size:16;not null"`
	IsLeaf  int    `gorm:"size:4;not null"`
	OrgName string `gorm:"size:32;not null"`
	Address string `gorm:"size:64;not null"`
	Phone   string `gorm:"size:13;not null"`
	Email   string `gorm:"size:32;not null"`
	Sort    int    `gorm:"size:4;not null"`
	Level   int    `gorm:"size:4;not null"`
	Status  int    `gorm:"size:4;not null"`
}
