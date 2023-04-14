package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"

	"github.com/spf13/viper"
	"github.com/tsunagatteru/ishiki-no-nagare/db"
	"github.com/tsunagatteru/ishiki-no-nagare/res"
	"github.com/tsunagatteru/ishiki-no-nagare/server"
)

func main() {
	var configPath string
	var resourcesPath string
	var dataPath string
	flag.StringVar(&configPath, "cfg", "/etc/inn/", "path to config")
	flag.StringVar(&resourcesPath, "res", "embed", "path to resources folder, can use embedded one")
	flag.StringVar(&dataPath, "data", "/var/lib/inn/", "path to store app data")
	flag.Parse()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	var resources fs.FS
	if resourcesPath == "embed" {
		resources = res.GetEmbedFS()
	} else {
		resources = os.DirFS(resourcesPath)
	}
	os.Mkdir(dataPath, 0755)
	os.Mkdir(dataPath+"images/", 0755)
	viper.Set("datapath", dataPath)
	dbConn := db.Open(dataPath)
	defer dbConn.Close()
	db.CreateTable(dbConn)
	server.RunRouter(dbConn, resources)
}
