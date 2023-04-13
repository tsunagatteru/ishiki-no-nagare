package config

import (
	"log"
	"os"

	"github.com/tsunagatteru/ishiki-no-nagare/model"
	"gopkg.in/yaml.v3"
)

func Read(path string) *model.Config {
	config := &model.Config{}
	file, err := os.Open(path)
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

func Write(cfg model.Config) {
}
