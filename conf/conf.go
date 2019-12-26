package conf

import "github.com/astaxie/beego"

var (
	Info, _ = beego.AppConfig.GetSection("info")
	Port = beego.AppConfig.String("httpport")
	Dmain = Info["domain_root"] + ":" + Port + "/"
)
