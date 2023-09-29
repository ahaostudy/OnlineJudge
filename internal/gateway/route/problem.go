package route

import (
	"github.com/gin-gonic/gin"

	"main/internal/gateway/controller/problem"
	"main/internal/gateway/middleware/jwt"
)

func RegisterProblemRouter(r *gin.RouterGroup) {
	r.GET("/", jwt.Parse(), problem.GetProblemList)
	r.GET("/:id", problem.GetProblem)
	r.GET("/count", problem.CreateProblemCount)
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
