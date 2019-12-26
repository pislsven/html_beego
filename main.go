package main

import (
	"github.com/astaxie/beego"
	"html_beego/conf"
	_ "html_beego/conf"
	_ "html_beego/routers"
	"strings"
)

func initStatis(){
	filePath := conf.Info["file_upload_collection"]
	filePath = strings.TrimSuffix(filePath, "/")
	filePathRouter := "/" + filePath

	beego.SetStaticPath(filePathRouter, filePath)
}

func main() {
	initStatis()
	beego.Run()
}

