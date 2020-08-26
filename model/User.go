package model

import (
	"context"
	"easy-admin/utils/respcode"
	"time"

	"github.com/jinzhu/gorm"
)

// User 用户表
type User struct {
	gorm.Model
	OrgID           uint       `gorm:"size:16;NOT NULL" json:"orgid"`
	Username        string     `gorm:"size:32;NOT NULL" json:"username"`
	Password        string     `gorm:"size:32;NOT NULL" json:"password"`
	Enabled         int        `gorm:"size:1;NOT NULL;default:1" json:"enable"`
	Locked          int        `gorm:"size:1;NOT NULL;default:0" json:"locked"`
	LockreleaseTime *time.Time `gorm:"size:16;type:timestamp" json:"lockreleasetime"`
}

func (u *User) clearPassword() {
	u.Password = ""
}

func clearPassword(s []User) {
	if len(s) == 0 {
		return
	}
	for u := range s {
		s[u].clearPassword()
	}
}

func userQueryByName(ctx context.Context, username string) (*User, int) {
	var user User
	db.Unscoped().Where("username = ?", username).First(&user)

	if user.ID > 0 {
		if user.Locked == 1 && user.Enabled == 0 {
			return &user, respcode.USER_LOCKED
		}
		return &user, respcode.USER_HAVED
	}
	return &user, respcode.USER_NOT_FIND
}
func userQueryByID(ctx context.Context, id uint) (*User, int) {
	var user User
	db.Unscoped().Where("id = ?", id).First(&user)

	if user.ID > 0 {
		if user.Locked == 1 && user.Enabled == 0 {
			return &user, respcode.USER_LOCKED
		}
		return &user, respcode.USER_HAVED
	}
	return &user, respcode.USER_NOT_FIND
}

func UserAdd(ctx context.Context, user *User) (*User, int) {
	defer user.clearPassword()
	_, code := userQueryByName(ctx, user.Username)

	if code == respcode.USER_NOT_FIND {
		db.Create(&user)
		return user, respcode.USER_CREATE_FINISHED
	}
	return nil, respcode.USER_HAVED
}

func UserDel(ctx context.Context, id uint) (*User, int) {
	user, code := userQueryByID(ctx, id)
	if code == respcode.USER_HAVED {
		user.Enabled = 0
		user.Locked = 1
		db.Save(&user)
		db.Delete(&user)
		return user, respcode.USER_LOCK_FINISHED
	}
	return user, code
}

func UserGet(ctx context.Context, id uint) (*User, int) {
	user, code := userQueryByID(ctx, id)
	defer user.clearPassword()
	return user, code
}

func UserList(ctx context.Context, user *User, CurPN uint) ([]User, int) {
	var users []User
	db.Unscoped().Offset((CurPN - 1) * 2).Limit(CurPN * 2).Where(user).Find(&users)

	defer clearPassword(users)
	return users, respcode.USER_FIND_FINISHED

}
