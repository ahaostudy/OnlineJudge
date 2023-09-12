package route

import (
	"main/internal/gateway/controller/problem"
	"main/internal/gateway/middleware/jwt"

	"github.com/gin-gonic/gin"
)

func RegisterProblemRouter(r *gin.RouterGroup) {
	r.GET("/", problem.GetProblemList)
	r.GET("/:id", problem.GetProblem)

	r.Use(jwt.Auth(), jwt.AuthAdmin())
	r.POST("/", problem.CreateProblem)
	r.PUT("/:id", problem.UpdateProblem)
	r.DELETE("/:id", problem.DeleteProblem)
}
