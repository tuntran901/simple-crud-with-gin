package main

import (
	"crud-note-simple/base"
	"crud-note-simple/server"
	"fmt"
)

func main() {
	fmt.Printf("CURD Simple Start")
	instance := server.StartAPISeverPort(base.APIPort)
	instance.Run()
}