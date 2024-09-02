package models

import (
	"fmt"
	"ginChat/utils"
	"time"

	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name          string
	PassWord      string
	Phone         string
	Email         string
	Identity      string
	ClientIP      string
	ClientPort    string
	LoginTime     time.Time `gorm:"default:null"`
	HeartbeatTime time.Time `gorm:"default:null"`
	LogoutTime    time.Time `gorm:"default:null"`
	IsLogout      bool
	DeviceInfo    string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

// 获取UserBasic数据库表的数据
func GetUserList() []*UserBasic {
	data := make([]*UserBasic, 10)
	utils.DB.Find(&data)
	for _, v := range data {
		fmt.Println(v)
	}
	return data
}

// 新增用户
func CreateUser(user UserBasic) {
	utils.DB.Create(&user)
}

// 删除用户
func DeleteUser(user UserBasic) {
	utils.DB.First(&user, user.ID)
	utils.DB.Delete(&user)
}

// 更新用户信息用户
func UpdateUser(user UserBasic) {
	utils.DB.First(&user, user.ID)
	utils.DB.Model(&user).Updates(
		UserBasic{
			Name:     user.Name,
			PassWord: user.PassWord,
			Phone:    user.Phone,
			Email:    user.Email,
		})
}
