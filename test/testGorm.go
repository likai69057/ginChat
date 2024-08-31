package main

import (
	"fmt"
	"ginChat/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:Alikai19940818@tcp(114.132.222.214:3306)/ginChat?charset=utf8mb4&parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.UserBasic{})

	user := &models.UserBasic{}
	user.Name = "申专"
	db.Create(user)

	fmt.Println(db.First(user, 1))
	db.Model(user).Update("PassWord", "1234")
}
