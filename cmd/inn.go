package main

import (
	"flag"
	"io/fs"
	"os"

	"github.com/tsunagatteru/ishiki-no-nagare/config"
	"github.com/tsunagatteru/ishiki-no-nagare/db"
	"github.com/tsunagatteru/ishiki-no-nagare/res"
	"github.com/tsunagatteru/ishiki-no-nagare/server"
)

func main() {
	var configPath string
	var resourcesPath string
	flag.StringVar(&configPath, "cfg", "/etc/inn/config.yml", "path to config")
	flag.StringVar(&resourcesPath, "res", "embed", "path to resources folder, can use embedded one")
	flag.Parse()
	config := config.Read(configPath)
	var resources fs.FS
	if resourcesPath == "embed" {
		resources = res.GetEmbedFS()
	} else {
		resources = os.DirFS(resourcesPath)
	}
	dataPath := config.DataPath
	os.Mkdir(dataPath, 0755)
	os.Mkdir(dataPath+"images/", 0755)
	dbConn := db.Open(dataPath)
	defer dbConn.Close()
	db.CreateTable(dbConn)
	router := server.NewRouter()
	server.RunRouter(router, dbConn, config, resources)
}
