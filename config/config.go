package config

import (
    "log"
    "os"

	"gopkg.in/yaml.v2"
	"github.com/tsunagatteru/ishiki-no-nagare/model"
)

func Read()(*model.Config){
	config := &model.Config{}
	file, err := os.Open("config.yml")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
    d := yaml.NewDecoder(file)
	if err := d.Decode(&config); err != nil {
        log.Fatalln(err)
    }
	return config
}
