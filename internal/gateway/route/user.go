package route

import (
	"main/internal/gateway/controller/user"
	"main/internal/gateway/middleware/jwt"

	"github.com/gin-gonic/gin"
)

func RegisterUserRouter(r *gin.RouterGroup) {
	r.POST("/register", user.Register)
	r.POST("/login", user.Login)
	r.POST("/captcha", user.GetCaptcha)

	r.Use(jwt.Auth())
	r.PUT("/:id", user.UpdateUser)

	r.Use(jwt.AuthAdmin())
	r.POST("/", user.CreateUser)
}
