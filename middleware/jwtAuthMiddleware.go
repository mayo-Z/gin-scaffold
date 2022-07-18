package middleware

import (
	"errors"
	"gin-scaffold/dao"
	"gin-scaffold/public"
	"github.com/gin-gonic/gin"
	"strings"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取Authorization
		tokenString := ctx.GetHeader("Authorization")
		//	valid token format
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			ResponseError(ctx, 2001, errors.New("权限不足"))
			return
		}
		tokenString = tokenString[7:]
		token, claims, err := public.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ResponseError(ctx, 2002, errors.New("权限不足"))
			return
		}
		//	验证通过后获取claim中的userid
		userId := claims.UserId
		DB := dao.GetDB()
		var user dao.Admin
		DB.First(&user, userId)
		if user.ID == 0 {
			ResponseError(ctx, 2003, errors.New("权限不足"))
			return
		}
		//	将用户信息写入上下文
		ctx.Set("user", user)
		ctx.Next()
	}
}
