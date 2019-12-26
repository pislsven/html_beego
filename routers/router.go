package routers

import (
	"html_beego/controllers"
	"github.com/astaxie/beego"
	"html_beego/controllers/editorctrller"
)

func init() {
    beego.Router("/", &controllers.MainController{})

	editorNS := beego.NewNamespace("/editor",
		beego.NSRouter("/k/index", &editorctrller.Kindeditorctrller{}, "*:Index"),
		beego.NSRouter("/k/upload", &editorctrller.Kindeditorctrller{}, "*:Upload"),
	)

	beego.AddNamespace(editorNS)
}
