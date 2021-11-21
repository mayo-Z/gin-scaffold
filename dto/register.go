package dto

import (
	"auth_frame/public"
	"github.com/gin-gonic/gin"
	"time"
)

type AdminSessionInfo struct {
	Uid       int       `json:"uid" `
	Username  string    `json:"username" `
	LoginTime time.Time `json:"login_time" `
}

type RegisterInput struct {
	Username string `json:"username" comment:"用户名" example:"admin" validate:"required"`
	Password string `gorm:"password" comment:"密码" example:"123456" validate:"required,valid_password"`
}

func (a *RegisterInput) BindValidParam(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, a)
}

type AdminLoginInput struct {
	Username string `json:"username" form:"username" comment:"用户名" example:"admin" validate:"required"`
	Password string `json:"password" comment:"密码" example:"123456" validate:"required,valid_password"`
}

func (a *AdminLoginInput) BindValidParam(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, a)
}

//type AdminLoginOutput struct {
//	Token string `json:"token" form:"token" comment:"token" example:"token" validate:""`
//}
