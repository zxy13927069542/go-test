package redis

import (
	"flag"
	log "github.com/pion/ion-log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"reflect"
)

var (
	path   = flag.String("path", "./config.yaml", "配置文件路径")
	config *Config
)

type Config struct {
	Redis struct {
		Addresses []string
		Pwd       string
	}
}

func Init() error {
	log.Infof("Loading Config...")
	if !reflect.DeepEqual(config, Config{}) {
		return nil
	}

	flag.Parse()
	file, err := ioutil.ReadFile(*path)
	if err != nil {
		log.Errorf("Read file Error! Reason: %v", err)
		return err
	}

	err = yaml.Unmarshal(file, config)
	if err != nil {
		log.Errorf("Unmarshal Config File Error! Reason: %v", err)
		return err
	}

	log.Infof("Load Config Success!")
	return nil
}
