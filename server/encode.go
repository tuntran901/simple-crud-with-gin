package server

import (
	"crud-note-simple/biz/dal/dataobject"
	"github.com/gin-gonic/gin"
	"net/http"
)

type APIResponse struct {
	Status		int32		 		`json:"status,omitempty"`
	Data		APIResponseData		`json:"data,omitempty"`
	Error		error				`json:"error,omitempty"`
}

type APIResponseData struct {
	Note		*dataobject.NotesDO		`json:"note,omitempty"`
	Notes		[]*dataobject.NotesDO	`json:"notes,omitempty"`
	NoteId		int64					`json:"note_id,omitempty"`
	IsUpdate	bool					`json:"is_update,omitempty"`
	IsDelete	bool					`json:"is_delete,omitempty"`
}

func StatusOK(c *gin.Context, r *APIResponseData)  {
	c.JSON(http.StatusOK, &APIResponse{
		Status: http.StatusOK,
		Data:   *r,
	})
}

func StatusBadRequest(c *gin.Context, err error)  {
	c.JSON(http.StatusBadRequest, &APIResponse{
		Status: http.StatusBadRequest,
		Error:   err,
	})
}
