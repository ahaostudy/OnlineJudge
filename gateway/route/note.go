package route

import (
	"main/gateway/controller/note"
	"main/gateway/middleware/jwt"

	"github.com/gin-gonic/gin"
)

func RegisterNoteRouter(r *gin.RouterGroup) {
	r.GET("/:id", note.GetNote)

	r.Use(jwt.Parse())
	r.GET("/", note.GetNoteList)

	r.Use(jwt.Auth())
	r.POST("/", note.CreateNote)
	r.PUT("/:id", note.UpdateNote)
	r.DELETE("/:id", note.DeleteNote)
}
