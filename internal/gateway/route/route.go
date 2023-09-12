package route

import (
	"main/internal/gateway/controller"
	"main/internal/gateway/middleware/jwt"

	"github.com/gin-gonic/gin"
)

func InitRoute() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/v1")
	RegisterUserRouter(api.Group("/user"))
	RegisterProblemRouter(api.Group("/problem"))

	api.Use(jwt.Auth())
	api.GET("/hello", controller.Hello)

	return r
}
