package controllers

import (
	"github.com/astaxie/beego"
	"myblog/tools"
)

type CategoryController struct {
	beego.Controller
}

//跳转到管理页面
func (c *CategoryController) Get() {
	c.Data["IsLogin"] = checkAccount(c.Ctx)

	category := c.Input().Get("category")
	if category == "" {
		c.TplName = "category.html"
		topics, err := tools.GetAllTopics()
		if err != nil {
			beego.Error(err)
		}
		c.Data["Topics"] = topics
	} else {
		topics, err := tools.GetCategoryContent(category)
		if err != nil {
			beego.Error(err)
		}
		//显示为分类的时候就不用页码了
		c.Data["Topics"] = topics
		c.TplName = "category.html"
	}

	var err error
	c.Data["Categories"], err = tools.GetCategoryName()
	if err != nil {
		beego.Error(err)
	}
}