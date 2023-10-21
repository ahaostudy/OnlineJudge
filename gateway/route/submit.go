package route

import (
	"github.com/gin-gonic/gin"

	"main/gateway/controller/submit"
	"main/gateway/middleware/jwt"
)

func RegisterSubmitRouter(r *gin.RouterGroup) {
	r.GET("/latest", jwt.Parse(), submit.GetLatestSubmits)
	r.GET("/calendar", submit.GetSubmitCalendar)

	r.Use(jwt.Auth())

	r.GET("/", submit.GetSubmitList)
	r.GET("/:id", submit.GetSubmit)
	r.POST("/", submit.Submit)
	r.DELETE("/:id", submit.DeleteSubmit)

	r.POST("/result", submit.GetResult)
	r.POST("/debug", submit.Debug)
}
