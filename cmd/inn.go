package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
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
	variables := viper.New()
	variables.Set("datapath", dataPath)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	viper.WatchConfig()
	var resources fs.FS
	if resourcesPath == "embed" {
		resources = res.GetEmbedFS()
	} else {
		resources = os.DirFS(resourcesPath)
	}
	os.Mkdir(dataPath, 0755)
	os.Mkdir(dataPath+"images/", 0755)
	viper.Set("datapath", dataPath)
	server.Run(resources, variables)
}
