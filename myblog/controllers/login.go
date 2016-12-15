package controllers

import (
	"github.com/astaxie/beego"
	_"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	//判断是否是退出登陆操作
	isExist := c.Input().Get("exit") == "true"
	if isExist {
		//目的为删除cookie
		c.Ctx.SetCookie("username", "", -1, "/")
		c.Ctx.SetCookie("password", "", -1, "/")

		//重定向到home
		c.Redirect("/", 302)
		return
	}

	c.TplName = "login.html"

}

//判断账号密码是否存在
func (c *LoginController) Post() {
	// 获取表单信息
	uname := c.Input().Get("username")
	pwd := c.Input().Get("password")

	if uname == "admin" && pwd == "admin" {
		beego.Error("账号密码正确！")

		//设置cookie
		autoLogin := c.Input().Get("autorLogin") == "on"
		maxAge := 0
		if autoLogin {
			// 这个代表的是一个星期
			//maxAge决定着Cookie的有效期，单位为秒
			maxAge = 7 * 86400
		}
		c.Ctx.SetCookie("username", uname, maxAge, "/")
		c.Ctx.SetCookie("password", pwd, maxAge, "/")

		//如果正确就登陆
		//跳转到首页
		c.Redirect("/", 302)
		return
	}
	beego.Error("账号密码错误！")
	//不正确就继续登陆
	c.Redirect("/login", 302)
}

func checkAccount(ctx *context.Context) bool {
	ck1, err := ctx.Request.Cookie("username")
	if err != nil {
		return false
	}
	uname := ck1.Value

	ck2, err := ctx.Request.Cookie("password")
	if err != nil {
		return false
	}
	pwd := ck2.Value

	return uname == ck1.Value && pwd == ck2.Value
}