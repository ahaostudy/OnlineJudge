package route

import (
	"main/internal/gateway/controller/problem"
	"main/internal/gateway/middleware/jwt"

	"github.com/gin-gonic/gin"
)

func RegisterProblemRouter(r *gin.RouterGroup) {
	r.GET("/", problem.GetProblemList)
	r.GET("/:id", problem.GetProblem)
	r.GET("/testcase/:id", problem.GetTestcase)

	r.Use(jwt.Auth(), jwt.AuthAdmin())

	r.POST("/", problem.CreateProblem)
	r.PUT("/:id", problem.UpdateProblem)
	r.DELETE("/:id", problem.DeleteProblem)
	r.POST("/testcase/", problem.CreateTestcase)
	r.DELETE("/testcase/:id", problem.DeleteTestcase)
	r.GET("/contest/:id", problem.GetContestProblem)
	r.GET("/contest/", problem.GetContestProblemList)
}
