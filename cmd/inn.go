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
	var dataPath string
	flag.StringVar(&configPath, "cfg", "/etc/inn/config.yml", "path to config")
	flag.StringVar(&resourcesPath, "res", "embed", "path to resources folder, can use embedded one")
	flag.StringVar(&dataPath, "data", "", "path to data, overwrites config")
	flag.Parse()
	config := config.Read(configPath)
	config.ConfigPath = configPath
	var resources fs.FS
	if resourcesPath == "embed" {
		resources = res.GetEmbedFS()
	} else {
		resources = os.DirFS(resourcesPath)
	}
	if dataPath == "" {
		dataPath = config.DataPath
	} else {
		config.DataPath = dataPath
	}
	os.Mkdir(dataPath, 0755)
	os.Mkdir(dataPath+"images/", 0755)
	dbConn := db.Open(dataPath)
	defer dbConn.Close()
	db.CreateTable(dbConn)
	server.RunRouter(dbConn, config, resources)
}
