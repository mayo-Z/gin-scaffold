package main

import (
	"auth_frame/dao"
	"auth_frame/route"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func main() {
	route.InitConfig()
	db := dao.InitDB()
	defer db.Close()
	r := route.InitRouter()
	port := viper.GetString("server.port")
	r.Run(":" + port)
}
