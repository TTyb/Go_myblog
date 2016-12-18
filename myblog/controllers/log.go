package controllers

import (
	"github.com/astaxie/beego"
	"myblog/tools"
)

type LogController struct {
	beego.Controller
}

func (c *LogController) Get() {

	//登陆状态就可以跳转
	c.TplName = "log.html"
	//检查是不是登陆状态
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.Data["Content"] = tools.GetLogContent()
}

func (c *LogController) Post() {
	logcontent := c.Input().Get("logcontent")
	if logcontent != "" {
		//写入日志
		err := tools.WriteLog(logcontent)
		if err != nil {
			beego.Error(err)
		}
	}
	c.Redirect("/log", 302)
}