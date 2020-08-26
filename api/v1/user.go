package v1

import (
	"easy-admin/model"
	"easy-admin/utils"
	"easy-admin/utils/respcode"

	"net/http"

	"github.com/gin-gonic/gin"
)

// UserAdd 添加用户
func UserAdd(c *gin.Context) {
	var user model.User
	ctx := c.Request.Context()
	c.ShouldBindJSON(&user)
	_, code := model.UserAdd(ctx, &user)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   user,
		"msg":    respcode.GetRespMsg(code),
	})
}

// UserMod 修改用户
func UserMod(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// UserDel 删除用户
func UserDel(c *gin.Context) {
	ctx := c.Request.Context()
	strid := c.Param("id")
	id := utils.StrToUint(strid)
	user, code := model.UserDel(ctx, id)
	c.JSON(200, gin.H{
		"status": code,
		"data":   user,
		"msg":    respcode.GetRespMsg(code),
	})
}

// UserInfo 查看用户信息
func UserInfo(c *gin.Context) {
	ctx := c.Request.Context()
	strid := c.Param("id")
	id := utils.StrToUint(strid)
	user, code := model.UserGet(ctx, id)
	c.JSON(200, gin.H{
		"status": code,
		"data":   user,
		"msg":    respcode.GetRespMsg(code),
	})
}

// UserList 获取用户列表
func UserList(c *gin.Context) {
	var user model.User
	ctx := c.Request.Context()
	strpage := c.Query("page")
	page := utils.StrToUint(strpage)
	c.ShouldBindJSON(&user)
	users, code := model.UserList(ctx, &user, page)
	c.JSON(200, gin.H{
		"status": code,
		"data":   users,
		"msg":    respcode.GetRespMsg(code),
	})
}
