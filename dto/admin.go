package dto

import (
	"gin-scaffold/public"
	"github.com/gin-gonic/gin"
	"time"
)

type AdminInfoOutput struct {
	Uid       int       `form:"uid"`
	Name      string    `form:"name"`
	LoginTime time.Time `form:"login_time" `
}

type ChangePwdInput struct {
	Password string `json:"password" form:"password" comment:"密码" example:"123456" validate:"required,valid_password"`
}

func (a *ChangePwdInput) BindValidParam(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, a)
}
