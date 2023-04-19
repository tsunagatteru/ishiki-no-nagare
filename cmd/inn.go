package main

import (
	"io/fs"
	"os"

	"github.com/tsunagatteru/ishiki-no-nagare/options"
	"github.com/tsunagatteru/ishiki-no-nagare/res"
	"github.com/tsunagatteru/ishiki-no-nagare/server"
)

func main() {
	variables := options.GetVar()
	resPath := variables.Get("res-path").(string)
	dataPath := variables.Get("data-path").(string)
	var resources fs.FS
	if resPath == "embed" {
		resources = res.GetEmbedFS()
	} else {
		resources = os.DirFS(resPath)
	}
	os.Mkdir(dataPath, 0755)
	os.Mkdir(dataPath+"images/", 0755)
	config := options.GetCfg(dataPath)
	server.Run(resources, variables, config)
}
