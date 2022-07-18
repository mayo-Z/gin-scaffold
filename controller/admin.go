package controller

import (
	"encoding/json"
	"fmt"
	"gin-scaffold/dao"
	"gin-scaffold/dto"
	"gin-scaffold/middleware"
	"gin-scaffold/public"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

type AdminController struct{}

func AdminRegister(group *gin.RouterGroup) {
	adminLogin := &AdminController{}
	group.GET("/info", adminLogin.AdminInfo)
	group.POST("/change_pwd", adminLogin.ChangePwd)
}

// AdminInfo godoc
// @Summary 用户信息
// @Description 用户信息
// @Tags 用户接口
// @ID /admin/info
// @Success 200 {object} middleware.Response{data=dto.AdminInfoOutput} "success"
// @Router /admin/info [get]
func (a *AdminController) AdminInfo(ctx *gin.Context) {
	//1.读取sessionKey对应json转换为结构体
	//2.取出数据然后封装输出结构体
	sess := sessions.Default(ctx)
	sessInfo := sess.Get(public.AdminSessionInfoKey)
	adminSessionInfo := &dto.AdminSessionInfo{}
	if err := json.Unmarshal([]byte(fmt.Sprint(sessInfo)),
		adminSessionInfo); err != nil {
		middleware.ResponseError(ctx, 2001, err)
		return
	}

	out := &dto.AdminInfoOutput{
		Uid:  adminSessionInfo.Uid,
		Name: adminSessionInfo.Username,
	}
	fmt.Println(out)
	middleware.ResponseSuccess(ctx, out)
}

// ChangePwd godoc
// @Summary 修改密码
// @Description 修改密码
// @Tags 用户接口
// @ID /admin/change_pwd
// @Accept  json
// @Produce  json
// @Param body body dto.ChangePwdInput true "body"
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /admin/change_pwd [post]
func (a *AdminController) ChangePwd(ctx *gin.Context) {
	params := &dto.ChangePwdInput{}
	if err := params.BindValidParam(ctx); err != nil {
		middleware.ResponseError(ctx, 2000, err)
		return
	}
	sess := sessions.Default(ctx)
	sessInfo := sess.Get(public.AdminSessionInfoKey)
	adminSessionInfo := &dto.AdminSessionInfo{}
	if err := json.Unmarshal([]byte(fmt.Sprint(sessInfo)),
		adminSessionInfo); err != nil {
		middleware.ResponseError(ctx, 2001, err)
		return
	}
	db := dao.GetDB()
	admin := &dao.Admin{}
	admin, err := admin.Find(db, adminSessionInfo.Uid)
	if err != nil {
		middleware.ResponseError(ctx, 2002, err)
		return
	}
	hashedPassword, _ := public.SetHashedPassword(params.Password)
	admin.Password = hashedPassword
	err = admin.Update(db)
	if err != nil {
		middleware.ResponseError(ctx, 2003, err)
		return
	}
	middleware.ResponseSuccess(ctx, "修改密码成功")
}
