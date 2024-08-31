package models

import (
	"fmt"
	"ginChat/utils"
	"time"

	"github.com/jinzhu/gorm"
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
