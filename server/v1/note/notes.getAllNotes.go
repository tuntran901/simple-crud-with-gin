package note

import (
	"crud-note-simple/server"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

// func GetAllNotes
func (s *NotesService) GetAllNotes(c *gin.Context) {
	currentPage, err := strconv.Atoi(c.Query("current_page"))
	if err != nil {
		fmt.Printf("GetAllNotes err: %+v", err)
		server.StatusBadRequest(c, err)
		return
	}

	if currentPage < 1 {
		currentPage = 1
	}

	perPage, err := strconv.Atoi(c.Query("per_page"))
	if err != nil {
		fmt.Printf("GetAllNotes err: %+v", err)
		server.StatusBadRequest(c, err)
		return
	}

	if perPage < 0 && perPage > 20 {
		perPage = 20
	}

	fmt.Printf("GetAllNotes reques: current_page=%d per_page=%d", currentPage, perPage)
	notes := s.NotesModel.GetAllNotes(int32(currentPage), int32(perPage))
	fmt.Printf("GetAllNotes Response: %+v", notes)

	res := &server.APIResponseData{Notes: notes}
	server.StatusOK(c, res)
	return
}
