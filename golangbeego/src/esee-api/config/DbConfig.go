package config

import (
	"fmt"
	"github.com/astaxie/beego/config"
)

var BConfig config.Configer

func init() {
	var err error
	BConfig, err = config.NewConfig("ini", "conf/jdbc.conf")
	if err != nil {
		fmt.Println("config init error:", err)
	}
}
