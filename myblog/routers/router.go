package routers

import (
	"myblog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/manage/topic", &controllers.ManageController{})
	// 这个是manage下面其他的东西
	beego.AutoRouter(&controllers.ManageController{})
	beego.Router("/topic", &controllers.TopicController{})
	//这个是views
	beego.AutoRouter(&controllers.TopicController{})
	//这个是log的
	beego.Router("/log", &controllers.LogController{})
	beego.AutoRouter(&controllers.LogController{})
	//错误机制
	beego.ErrorController(&controllers.ErrorController{})
}
