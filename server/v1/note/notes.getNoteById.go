package note

import (
	"crud-note-simple/server"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

// func GetNoteById
func(s *NotesService) GetNoteById(c *gin.Context)  {
	noteId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		fmt.Printf("GetNoteById errr: %+v", err)
		server.StatusBadRequest(c, err)
		return
	}

	fmt.Printf("GetNoteById reques: %+v", noteId)
	note := s.NotesModel.GetNoteById(noteId)
	fmt.Printf("GetNoteById Response: %+v", note)

	res := &server.APIResponseData{Note: note}
	server.StatusOK(c, res)
	return
}