package middleware

import (
	"mymodule/common"
	"mymodule/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authmiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取authorization header
		tokenString := ctx.GetHeader("Authorization")
		//验证格式：Authorization:Bearer tokenString
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			ctx.Abort()
			return
		}
		//验证token
		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			ctx.Abort()
			return
		}
		//验证通过后获取claims中的userId
		userid := claims.UserId
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userid)
		//用户
		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			ctx.Abort()
			return
		}
		//将userId写入上下文
		ctx.Set("user", user)
		ctx.Next()
	}
}
