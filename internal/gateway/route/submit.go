package route

import (
	"main/internal/gateway/controller/submit"
	"main/internal/gateway/middleware/jwt"

	"github.com/gin-gonic/gin"
)

func RegisterSubmitRouter(r *gin.RouterGroup) {
	r.Use(jwt.Auth())

	r.POST("/", submit.GetSubmitList)
	r.GET("/:id", submit.GetSubmit)
	r.DELETE("/:id", submit.DeleteSubmit)

	r.POST("/judge", submit.Submit)
	r.POST("/result", submit.GetResult)
	r.POST("/debug", submit.Debug)
}
