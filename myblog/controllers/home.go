package controllers

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {

	//登陆状态就可以跳转
	c.TplName = "home.html"
	//检查是不是登陆状态
	c.Data["IsLogin"] = checkAccount(c.Ctx)
}
