package main

import (
	"embed"

	"github.com/tsunagatteru/ishiki-no-nagare/config"
	"github.com/tsunagatteru/ishiki-no-nagare/db"
	"github.com/tsunagatteru/ishiki-no-nagare/server"
)

//go:embed res
var resources embed.FS

func main() {
	config := config.Read()
	dbConn := db.Open()
	defer dbConn.Close()
	db.CreateTable(dbConn)
	router := server.NewRouter()
	server.RunRouter(router, dbConn, config, resources)
}
