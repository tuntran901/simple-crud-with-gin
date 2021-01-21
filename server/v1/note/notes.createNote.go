package note

import (
	"crud-note-simple/biz/dal/dataobject"
	"crud-note-simple/server"
	"fmt"
	"github.com/gin-gonic/gin"
)

// func GetNoteById
func(s *NotesService) CreateNote(c *gin.Context)  {
	var note dataobject.NotesDO
	err := c.BindJSON(&note)
	if err != nil {
		fmt.Printf("CreateNote err: %+v", err)
		server.StatusBadRequest(c, err)
		return
	}

	if note.Content == "" || note.Title == "" {
		server.StatusBadRequest(c, fmt.Errorf("title is empty"))
		return
	}

	fmt.Printf("CreateNote reques: %+v", note)
	noteId := s.NotesModel.CreateNote(note.Title, note.Content)
	fmt.Printf("CreateNote Response: %+v", noteId)

	res := &server.APIResponseData{NoteId: noteId}
	server.StatusOK(c, res)
	return
}