package data

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Data struct {
	FastHands map[string]map[string]interface{} `yaml:"FastHands"`
	MetaData  map[string]string                 `yaml:"MetaData"`
	Workouts  map[string]map[string]interface{} `yaml:"Workouts"`
	Drinks    map[string]map[string]interface{} `yaml:"Drinks"`
	Misic     map[string]map[string]string      `yaml:"Misic"`
}

func Read() (Data, error) {
	var data Data
	yamlFile, err := os.ReadFile("./config/data.yaml")
	if err != nil {
		return data, fmt.Errorf("error reading config file: %v", err)
	}
	err = yaml.Unmarshal(yamlFile, &data)
	if err != nil {
		return data, fmt.Errorf("error unmarshaling yaml: %v", err)
	}
	return data, nil
}
