package route

import (
	"main/internal/gateway/controller/submit"
	"main/internal/gateway/middleware/jwt"

	"github.com/gin-gonic/gin"
)

func RegisterSubmitRouter(r *gin.RouterGroup) {
	r.Use(jwt.Auth())

	r.POST("/judge", submit.Submit)
	r.POST("/result", submit.GetResult)
}
