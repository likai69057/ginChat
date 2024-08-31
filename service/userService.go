package service

import (
	"ginChat/models"

	"github.com/gin-gonic/gin"
)

// GetUserList
// @Tags 用户列表
// @Success 200 {string} json{"code", "message"}
// @Router /user/getUserList [get]
func GetUserList(ctx *gin.Context) {
	data := models.GetUserList()

	ctx.JSON(200, gin.H{
		"message": data,
	})
}
