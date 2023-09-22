package route

import (
	"main/internal/gateway/controller/contest"
	"main/internal/gateway/middleware/jwt"

	"github.com/gin-gonic/gin"
)

func RegisterContestRouter(r *gin.RouterGroup) {
	r.GET("/", contest.GetContestList)
	r.POST("/rank", contest.ContestRank)

	r.Use(jwt.Auth())
	r.GET("/:id", contest.GetContest)
	r.POST("/register", contest.RegisterContest)

	r.Use(jwt.AuthAdmin())
	r.POST("/", contest.CreateContest)
	r.PUT("/:id", contest.UpdateContest)
	r.DELETE("/:id", contest.DeleteContest)
}
