package route

import (
	"github.com/gin-gonic/gin"

	"main/gateway/controller/user"
	"main/gateway/middleware/jwt"
)

func RegisterUserRouter(r *gin.RouterGroup) {
	r.POST("/register", user.Register)
	r.POST("/login", user.Login)
	r.POST("/captcha", user.GetCaptcha)
	r.GET("/avatar/:avatar", user.GetAvatar)

	r.Use(jwt.Parse())
	r.GET("/:id", user.GetUser)

	r.Use(jwt.Auth())
	r.PUT("/:id", user.UpdateUser)
	r.PUT("/avatar", user.UpdateAvatar)
	r.DELETE("/avatar", user.DeleteAvatar)

	r.Use(jwt.AuthAdmin())
	r.POST("/", user.CreateUser)
}
