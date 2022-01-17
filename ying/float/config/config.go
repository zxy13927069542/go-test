package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Node struct {
	F float64 `yaml:"f"`
	N int     `yaml:"n"`
}

type Config struct {
	Nodes []Node `yaml:"nodes"`
}

func GetConfig() *Config {
	file := "config/test.yaml"
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("read config file %s fail, reason: %s\n", file, err)
		return nil
	}

	var conf Config
	err = yaml.Unmarshal(bytes, &conf)
	if err != nil {
		fmt.Printf("ummarshal config %s fail, reason: %s\n", file, err)
		return nil
	}

	return &conf
}
