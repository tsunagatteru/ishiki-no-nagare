package options

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func GetVar() *viper.Viper {
	flag.String("res-path", "embed", "path to resources folder, can use embedded one")
	flag.String("data-path", "/app/data/", "path to store app data")
	flag.Int("page-length", 20, "number of post fetched per page request")
	flag.Parse()
	variables := viper.New()
	variables.BindPFlags(flag.CommandLine)
	return variables
}

func GetCfg(dataPath string) *viper.Viper {
	config := viper.New()
	config.SetDefault("host", "0.0.0.0")
	config.SetDefault("port", "8080")
	config.SetDefault("username", "admin")
	config.SetDefault("password", "admin")
	config.SetDefault("cookiekey", "secret")
	config.SetConfigName("config")
	config.SetConfigType("yaml")
	config.AddConfigPath(dataPath)
	err := config.ReadInConfig()
	if err != nil {
		log.Fatalln("error reading config file: %w", err)
	}
	config.WriteConfig()
	config.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	config.WatchConfig()
	return config
}
