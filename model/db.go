package model

import (
	"easy-admin/utils"
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	// 导入mysql驱动包
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(utils.DbDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		utils.DbUser, utils.DbPass, utils.DbHost, utils.DbPort, utils.DbName))
	if err != nil {
		fmt.Println("database 初始化错误1")
		os.Exit(1)
	}

	if db.DB().Ping() != nil {
		fmt.Println("database 初始化错误2")
		os.Exit(1)
	}

	// defer db.Close()
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "sys_" + defaultTableName
	}
	db.SingularTable(true)
	db.AutoMigrate(&User{}, &Menu{}, &Org{}, &Role{}, &RoleMenu{}, &UserRole{})
	db.DB().SetMaxIdleConns(5)
	db.DB().SetMaxOpenConns(10)
	db.DB().SetConnMaxLifetime(10 * time.Second)
}
