package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

type DbConfig struct {
	DbType string `yaml:"DbType"`
	User string `yaml:"User"`
	Password string `yaml:"Password"`
	Host string `yaml:"Host"`
	Port int `yaml:"Port"`
	DbName string `yaml:"DbName"`
	MaxIdleConns int `yaml:"MaxIdleConns"`
	MaxOpenConns int `yaml:"MaxOpenConns"`
	MaxLifetime int `yaml:"MaxLifetime"`	// s
}

type Config struct {
	Prefix string `yaml:"prefix"`
	ProxyUrl string `yaml:"proxyUrl"`
	Port int `yaml:"port"`
	DbConfig *DbConfig `yaml:"DbConfig"`
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

	// 打印调试
	//logger.Log.Debug(*c)
	//logger.Log.Debug(c.DbConfig)
	//setting, _ := json.MarshalIndent(c, "", "\t")
	//logger.Log.Debug(string(setting))

	return c
}