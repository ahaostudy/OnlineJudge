package route

import (
	"main/server/controller"
	"main/server/middleware/jwt"

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
