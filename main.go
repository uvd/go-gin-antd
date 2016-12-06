package main

import "github.com/urfave/negroni"
import "antd-admin-golang/config"

import "log"

func main() {
	//默认
	n := negroni.Classic()

	config := config.New()
	//初始化路由请求
	n.UseHandler(initRouter())

	if config.Web.Port == "" {
		log.Println("Config web Port is null ,set default is 9999")
		config.Web.Port = ":9999"
	} else {
		config.Web.Port = ":" + config.Web.Port
	}

	n.Run(config.Web.Port)
}
