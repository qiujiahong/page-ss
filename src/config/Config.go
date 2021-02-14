package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

type Config struct {
	ProxyUrl string `yaml:"proxyUrl"`
	Port int `yaml:"port"`
}

var Global = &Config{}

func Init () {
	Global.getConf()
}


func (c *Config) getConf() *Config {
	yamlFile, err := ioutil.ReadFile("conf/conf.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}