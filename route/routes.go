package route

import (
	"gin-scaffold/controller"
	"gin-scaffold/dao"
	"gin-scaffold/docs"
	"gin-scaffold/middleware"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"log"
	"os"
)

// @title Swagger Example API
// @version 1.0
// @description my blog

// @host localhost:8090
// @BasePath /api

func InitConfig() {
	docs.SwaggerInfo.Title = viper.GetString("swagger.title")
	docs.SwaggerInfo.Description = viper.GetString("swagger.desc")
	docs.SwaggerInfo.Host = viper.GetString("swagger.host")
	docs.SwaggerInfo.BasePath = viper.GetString("swagger.base_path")

	work, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(work + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic("err")
	}
}

func InitRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	router := gin.Default()
	router.Use(middlewares...)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//非登陆接口-----------------------------------------------------------------
	adminLoginRouter := router.Group("/auth")
	store, err := dao.GetSessionStore()
	if err != nil {
		log.Fatalf("sessions.NewRedisStoreerr:%v", err)
	}

	adminLoginRouter.Use(
		sessions.Sessions("mySession", store),
		middleware.TranslationMiddleware())
	{
		controller.AdminLoginRegister(adminLoginRouter)
	}
	//用户接口-----------------------------------------------------------------
	adminRouter := router.Group("/admin")
	adminRouter.Use(sessions.Sessions("mySession", store),
		middleware.SessionAuthMiddleware(),
		middleware.TranslationMiddleware())
	{
		controller.AdminRegister(adminRouter)
	}

	//router.Static("/dist", "./dist")
	return router
}
