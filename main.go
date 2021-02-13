package main

//# www.baidu.com/render     代理到 www.baidu.com     截图
//# www.baidu.com/render/abc 代理到 www.baidu.com/abc 截图
//proxyUrl: www.baidu.com
//# 是否转发header
//forwardHeader: false
//# 是否转发cookie
//forwardCookie: false


import (
	"fmt"
	_ "github.com/icattlecoder/godaemon"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

type conf struct {
	ProxyUrl string `yaml:"proxyUrl"`
	ForwardHeader string `yaml:"forwardHeader"`
	ForwardCookie string `yaml:"forwardCookie"`
}

func (c *conf) getConf() *conf {
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

func main() {
	var c conf
	conf :=	c.getConf()
	fmt.Printf("config:%v",conf)
}
