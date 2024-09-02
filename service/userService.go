package service

import (
	"ginChat/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetUserList
// @Summary 所有用户
// @Tags 用户模块
// @Success 200 {string} json{"code", "message"}
// @Router /user/getUserList [get]
func GetUserList(ctx *gin.Context) {
	data := models.GetUserList()

	ctx.JSON(200, gin.H{
		"message": data,
	})
}

// CreateUser
// @Summary 新增用户
// @Tags 用户模块
// @param name query string false "用户名"
// @param password query string false "密码"
// @param repassword query string false "确认密码"
// @Success 200 {string} json{"code", "message"}
// @Router /user/createUser [get]
func CreateUser(ctx *gin.Context) {
	user := models.UserBasic{}
	user.Name = ctx.Query("name")
	passWord := ctx.Query("password")
	rePassWord := ctx.Query("repassword")
	if passWord != rePassWord {
		ctx.JSON(-1, gin.H{
			"message": "密码不一致",
		})
		return
	}
	user.PassWord = passWord
	models.CreateUser(user)
	ctx.JSON(200, gin.H{
		"message": "新增用户成功！",
	})
}

// DeleteUser
// @Summary 删除用户
// @Tags 用户模块
// @param ID query string false "删除用户ID"
// @Success 200 {string} json{"code", "message"}
// @Router /user/deleteUser [get]
func DeleteUser(ctx *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(ctx.Query("ID"))
	user.ID = uint(id)
	models.DeleteUser(user)
	ctx.JSON(200, gin.H{
		"message": "删除用户成功！",
	})
}

// UpdateUser
// @Summary 更新用户
// @Tags 用户模块
// @param ID formData string false "用户ID"
// @param name formData string false "用户名"
// @param password formData string false "密码"
// @param phone formData string false "电话"
// @param email formData string false "邮箱"
// @Success 200 {string} json{"code", "message"}
// @Router /user/updateUser [post]
func UpdateUser(ctx *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(ctx.PostForm("ID"))
	user.ID = uint(id)
	user.Name = ctx.PostForm("name")
	user.PassWord = ctx.PostForm("password")
	user.Phone = ctx.PostForm("phone")
	user.Email = ctx.PostForm("email")
	models.UpdateUser(user)
	ctx.JSON(200, gin.H{
		"message":  "修改用户成功！",
		"name":     user.Name,
		"password": user.PassWord,
	})
}
