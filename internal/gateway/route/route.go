package route

import (
	"main/internal/gateway/controller/chatgpt"
	"main/internal/gateway/middleware/cors"
	"main/internal/gateway/middleware/jwt"

	"github.com/gin-gonic/gin"
)

func InitRoute() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/v1")
	api.Use(cors.Cors())

	RegisterUserRouter(api.Group("/user"))
	RegisterProblemRouter(api.Group("/problem"))
	RegisterSubmitRouter(api.Group("/submit"))

	api.Use(jwt.Auth())
	api.POST("/chat", chatgpt.Chat)

	return r
}
