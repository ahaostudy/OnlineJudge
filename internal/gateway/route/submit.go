package route

import (
	"main/internal/gateway/controller/submit"
	"main/internal/gateway/middleware/jwt"

	"github.com/gin-gonic/gin"
)

func RegisterSubmitRouter(r *gin.RouterGroup) {
	r.Use(jwt.Auth())

	r.GET("/", submit.GetSubmitList)
	r.GET("/:id", submit.GetSubmit)
	r.POST("/", submit.Submit)
	r.DELETE("/:id", submit.DeleteSubmit)

	r.POST("/result", submit.GetResult)
	r.POST("/debug", submit.Debug)
}
