package route

import (
	"main/internal/gateway/controller/testcase"
	"main/internal/gateway/middleware/jwt"

	"github.com/gin-gonic/gin"
)

func RegisterTestcaseRouter(r *gin.RouterGroup) {
	r.GET("/:id", testcase.GetTestcase)

	r.Use(jwt.Auth(), jwt.AuthAdmin())

	r.POST("/", testcase.CreateTestcase)
	r.DELETE("/:id", testcase.DeleteTestcase)
}
