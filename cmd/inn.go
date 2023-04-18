package main

import (
	"io/fs"
	"os"

	"github.com/spf13/viper"
	"github.com/tsunagatteru/ishiki-no-nagare/options"
	"github.com/tsunagatteru/ishiki-no-nagare/res"
	"github.com/tsunagatteru/ishiki-no-nagare/server"
)

func main() {
	variables, config := options.Get()
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
	viper.Set("datapath", dataPath)
	server.Run(resources, variables, config)
}
