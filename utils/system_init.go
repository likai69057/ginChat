package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// 用于解析json配置文件
func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println("config app:", viper.GetString("dsn"))
	//fmt.Println("config app:", viper.Get("mysql"))
}

// 通过配置文件初始化数据库表
func Initmysql() {

	//自定义日志模板 sql日志打印
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, //慢sql阈值
			LogLevel:      logger.Info, //日志级别
			Colorful:      true,        //彩色
		},
	)

	DB, _ = gorm.Open(mysql.Open(viper.GetString("dsn")), &gorm.Config{Logger: newLogger})

}
