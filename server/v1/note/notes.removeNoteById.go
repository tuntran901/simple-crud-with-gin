package note

import (
	"crud-note-simple/server"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

// func RemoveNoteById
func(s *NotesService) RemoveNoteById(c *gin.Context)  {
	noteId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		fmt.Printf("RemoveNoteById errr: %+v", err)
		server.StatusBadRequest(c, err)
		return
	}

	fmt.Printf("RemoveNoteById reques: %+v", noteId)
	isDelete := s.NotesModel.DeleteNote(noteId)
	fmt.Printf("RemoveNoteById Response: %+v", isDelete)

	res := &server.APIResponseData{IsDelete: isDelete}
	server.StatusOK(c, res)
	return

}