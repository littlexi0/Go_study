package common

import (
	"fmt"
	"mymodule/model"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// driverName := "mysql"
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		database,
		charset)

	var err error
	DB, err = gorm.Open(mysql.Open(args), &gorm.Config{})
	// db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to connect database,err: " + err.Error())
	}
	DB.AutoMigrate(&model.User{})
}

func GetDB() *gorm.DB {
	return DB
}
