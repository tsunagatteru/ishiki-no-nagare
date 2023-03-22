package main

import (
	"embed"
	"flag"
	"io/fs"
	"log"
	"os"

	"github.com/tsunagatteru/ishiki-no-nagare/config"
	"github.com/tsunagatteru/ishiki-no-nagare/db"
	"github.com/tsunagatteru/ishiki-no-nagare/server"
)

//go:embed res
var embedFS embed.FS

func main() {
	var configPath string
	var resourcesPath string
	flag.StringVar(&configPath, "cfg", "config.yml", "path to config")
	flag.StringVar(&resourcesPath, "res", "embed", "path to resources folder, can use embedded one")
	flag.Parse()
	log.Println(resourcesPath)
	config := config.Read(configPath)
	var resources fs.FS
	if resourcesPath == "embed" {
		resRoot, err := fs.Sub(embedFS, "res")
		if err != nil {
			log.Fatalln(err)
		}
		resources = resRoot
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
