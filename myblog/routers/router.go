package routers

import (
	"myblog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/manage", &controllers.ManageController{})
	// 这个是topic_add
	beego.AutoRouter(&controllers.ManageController{})
}
