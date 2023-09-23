package jwt

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"main/api/user"
	"main/internal/common/code"
	"main/rpc"
)

// Auth JWT鉴权中间件
func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		prefix := "Bearer "

		// 获取token，如果token格式不合法则拦截
		token := ctx.GetHeader("Authorization")
		if !strings.HasPrefix(token, prefix) {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// 去掉多余的prefix
		token = strings.TrimPrefix(token, prefix)

		// 解析并校验Token
		id, ok := ParseToken(token)
		if !ok {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.Set("user_id", id)
		ctx.Next()
	}
}

// AuthAdmin 管理员身份鉴权
// 使用该中间件前需要先经过Auth中间件
func AuthAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.GetInt64("user_id")
		result, err := rpc.UserCli.IsAdmin(ctx.Request.Context(), &rpcUser.IsAdminRequest{ID: id})
		if err != nil || result.GetStatusCode() != code.CodeSuccess.Code() {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		if !result.GetIsAdmin() {
			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}

		ctx.Set("is_admin", true)
		ctx.Next()
	}
}
