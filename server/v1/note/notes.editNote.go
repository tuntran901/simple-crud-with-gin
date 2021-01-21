package note

import (
	"crud-note-simple/biz/dal/dataobject"
	"crud-note-simple/server"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

// func EditNote
func (s *NotesService) EditNote(c *gin.Context) {
	note := &dataobject.NotesDO{}
	err := c.BindJSON(note)
	if err != nil {
		fmt.Printf("EditNote errr: %+v", err)
		server.StatusBadRequest(c, err)
		return
	}

	note.Id, err = strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		fmt.Printf("EditNote errr: %+v", err)
		server.StatusBadRequest(c, err)
		return
	}

	fmt.Printf("EditNote reques: %+v", note)
	isUpdate := s.NotesModel.UpdateNote(note)
	fmt.Printf("EditNote Response: %+v", isUpdate)

	res := &server.APIResponseData{IsUpdate: isUpdate}
	server.StatusOK(c, res)
	return
}
