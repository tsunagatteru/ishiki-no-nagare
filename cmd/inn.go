package main

import (
	"io/fs"
	"log"
	"os"

	"github.com/tsunagatteru/ishiki-no-nagare/options"
	"github.com/tsunagatteru/ishiki-no-nagare/res"
	"github.com/tsunagatteru/ishiki-no-nagare/server"
)

func main() {
	variables := options.GetVar()
	resPath := variables.GetString("res-path")
	dataPath := variables.GetString("data-path")
	var resources fs.FS
	if resPath == "embed" {
		resources = res.GetEmbedFS()
	} else {
		resources = os.DirFS(resPath)
	}
	os.Mkdir(dataPath, 0755)
	os.Mkdir(dataPath+"images/", 0755)
	if _, err := os.Stat(dataPath + "config.yml"); os.IsNotExist(err) {
		file, err := os.Create(dataPath + "config.yml")
		if err != nil {
			log.Fatalln("Unable to create file:", err)
		}
		defer file.Close()
	}
	config := options.GetCfg(dataPath)
	server.Run(resources, variables, config)
}
