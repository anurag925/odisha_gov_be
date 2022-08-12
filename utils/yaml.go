package utils

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

func LoadYaml(path string, data *map[any]any) {
	yfile, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}
	err2 := yaml.Unmarshal(yfile, &data)
	if err2 != nil {

		log.Fatal(err2)
	}
}
