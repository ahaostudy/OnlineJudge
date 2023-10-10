package route

import (
	"os"

	"github.com/gin-gonic/gin"

	"main/gateway/config"
	"main/gateway/controller/chatgpt"
	"main/gateway/middleware/cors"
	"main/gateway/middleware/jwt"
)

func InitRoute() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Cors())

	initStatic()
	r.Static(config.Config.Static.URI, config.Config.Static.Path)

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, "pong")
	})

	api := r.Group("/api/v1")

	RegisterUserRouter(api.Group("/user"))
	RegisterProblemRouter(api.Group("/problem"))
	RegisterSubmitRouter(api.Group("/submit"))
	RegisterContestRouter(api.Group("/contest"))

	api.Use(jwt.Auth())
	api.POST("/chat", chatgpt.Chat)

	return r
}

func initStatic() {
	if err :=os.MkdirAll(config.Config.Static.Path, os.ModePerm); err != nil {
		panic(err)
	}
}
