package options

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func Get() (*viper.Viper, *viper.Viper) {
	flag.String("config-path", "/etc/inn/", "path to config")
	flag.String("res-path", "embed", "path to resources folder, can use embedded one")
	flag.String("data-path", "/var/lib/inn/", "path to store app data")
	flag.Int("page-length", 20, "number of post fetched per page request")
	flag.Parse()
	variables := viper.New()
	variables.BindPFlags(flag.CommandLine)
	config := viper.New()
	config.SetConfigName("config")
	config.SetConfigType("yaml")
	config.AddConfigPath(variables.Get("config-path").(string))
	err := config.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	config.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	config.WatchConfig()
	return variables, config
}
