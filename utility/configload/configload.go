package configload 

import (
  "io/ioutil"
  "gopkg.in/yaml.v2"
)

type Conf struct {
  ServicePort string  `yaml:"service_port"`
}


func LoadConfig(filename string) (*Conf, error) {
	// Read the YAML file
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Create a new Conf instance
	config := &Conf{}

	// Unmarshal the YAML data into the Conf struct
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
