package main

import (
	//"github.com/gin-gonic/gin"
	"github.com/DeniesKresna/gocms/Config"
	"github.com/DeniesKresna/gocms/Routers"
)

var err error

func main() {
	Config.ConnectDB()

	r := Routers.SetupRouter()
	// running
	r.Run()
}
