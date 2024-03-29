package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var db *gorm.DB
var err error

func InitDB() {
	driverName := viper.GetString("datasource.driverName")
	name := viper.GetString("datasource.name")
	password := viper.GetString("datasource.password")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	charset := viper.GetString("datasource.charset")

	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&&parseTime=true",
		name,
		password,
		host,
		port,
		database,
		charset)

	db, err = gorm.Open(driverName, args)
	if err != nil {
		panic("failed to connect sql,err:" + err.Error())
	}
	db.AutoMigrate(&Admin{})
}

func GetDB() *gorm.DB {
	return db
}
