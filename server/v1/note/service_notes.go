package note

import (
	"crud-note-simple/base"
	"crud-note-simple/biz/core/notes"
	"github.com/gin-gonic/gin"
)

type NotesService struct {
	Route 	*gin.RouterGroup

	NotesModel *notes.NotesModel
}

func RegisterNotesService(r *gin.RouterGroup) (service *NotesService) {
	// init notes
	service = &NotesService{}
	service.Route = r

	// Register Model
	service.NotesModel = notes.NewNotesModel()

	// Install
	service.InstallRoute()
	service.InstallModel()
	return service
}

func (s *NotesService) InstallModel(){
	db := base.ConnectMysql()
	s.NotesModel.InstallModel(db)
}

func (s *NotesService) InstallRoute(){
	s.Route.Group("/notes")

	s.Route.POST("",s.CreateNote)
	s.Route.GET("",s.GetAllNotes)

	const noteId = "/:id"
	s.Route.GET(noteId,s.GetNoteById)
	s.Route.PUT(noteId,s.EditNote)
	s.Route.DELETE(noteId,s.RemoveNoteById)
}

