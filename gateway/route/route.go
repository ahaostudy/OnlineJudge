package route

import (
	"github.com/gin-gonic/gin"

	"main/gateway/controller/chatgpt"
	"main/gateway/middleware/cors"
	"main/gateway/middleware/jwt"
)

func InitRoute() *gin.Engine {
	r := gin.Default()

	// TODO: static path
	r.Static("static", "data/static")

	api := r.Group("/api/v1")
	api.Use(cors.Cors())

	RegisterUserRouter(api.Group("/user"))
	RegisterProblemRouter(api.Group("/problem"))
	RegisterSubmitRouter(api.Group("/submit"))
	RegisterContestRouter(api.Group("/contest"))

	api.Use(jwt.Auth())
	api.POST("/chat", chatgpt.Chat)

	return r
}
