package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"time"
	"myblog/tools"
)

type ManageController struct {
	beego.Controller
}

//跳转到管理页面
func (c *ManageController) Get() {
	if checkAccount(c.Ctx) == true {
		c.Data["IsLogin"] = true
		c.TplName = "manage.html"
	} else {
		c.TplName = "home.html"
	}
}

// 跳转到topic_add页面
func (c *ManageController) Add() {
	//检查是不是登陆状态
	if checkAccount(c.Ctx) == true {
		c.Data["IsLogin"] = true
		//登陆状态就可以跳转
		c.TplName = "topic_add.html"
	} else {
		//不是登陆状态就老老实实的去登陆
		c.TplName = "home.html"
	}
}

// 添加文章
func (c *ManageController) Post() {
	if !checkAccount(c.Ctx) {
		c.Redirect("/", 302)
		return
	}

	// 解析表单
	title := c.Input().Get("title")
	content := c.Input().Get("content")
	category := c.Input().Get("category")

	fmt.Println(title, content, category)


	//下面是将其保存下来

	//创建文件夹，这里是创建分类
	tools.Createdir("topic", category)
	//当前时间
	nowtime := time.Now().Format("2006-01-02")
	//写入内容到md
	mdname := nowtime + "#" + title + ".md"
	tools.Writefile(category, mdname, content)

	c.Redirect("/manage", 302)
}

