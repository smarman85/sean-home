package data

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Data struct {
	FastHands map[string]map[string]interface{} `yaml:"FastHands"`
	MetaData  map[string]string                 `yaml:"MetaData"`
	Weights   map[string]map[string]interface{} `yaml:"Weights"`
	Drinks    map[string]map[string]interface{} `yaml:"Drinks"`
	Misic     map[string]map[string]string      `yaml:"Misic"`
	Cardio    map[string]map[string]interface{} `yaml:"Cardio"`
}

func Read() (Data, error) {
	fmt.Println("Hello from data")
	var data Data
	yamlFile, err := ioutil.ReadFile("./config/data.yaml")
	if err != nil {
		return data, fmt.Errorf("error reading yaml file %v", err)
	}

	err = yaml.Unmarshal(yamlFile, &data)
	if err != nil {
		return data, fmt.Errorf("can't unmarshal yaml %v", err)
	}
	return data, nil
}
