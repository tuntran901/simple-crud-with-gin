package server

import (
	v1 "crud-note-simple/server/v1"
	"github.com/gin-gonic/gin"
)

type ApiSever struct {
	Port			string
	Routes 			*gin.Engine
}

func StartAPISeverPort(port string) *ApiSever {
	return &ApiSever{Port: port}
}

func (a *ApiSever) Run() {
	a.Routes = gin.Default()

	version1 := a.Routes.Group("v1")
	v1.InitializationRoute(version1)

	a.Routes.Run(a.Port)
}