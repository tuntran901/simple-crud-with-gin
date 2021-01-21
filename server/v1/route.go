package v1

import (
	"crud-note-simple/server/v1/note"
	"github.com/gin-gonic/gin"
)

func InitializationRoute(r *gin.RouterGroup) {
	note.RegisterNotesService(r)
}