package config

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)


func TestConfigFile(t *testing.T) {
	Convey("Init config file", t, func() {
		Init()
		fmt.Printf("Global=%v",Global)
		pwd, _ := os.Getwd()
		fmt.Printf("pwd=%v",pwd)
		Convey("When 1 integer is incremented", func() {
			So(Global,ShouldNotBeNil)
			So(Global.ProxyUrl,ShouldNotBeNil)
			//So(Global,ShouldBeNil)s
		})
	})
}