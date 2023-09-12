package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status_code": 200,
		"status_msg":  "OK",
		"data":        fmt.Sprintf("hello %d", c.GetInt64("user_id")),
	})
}
