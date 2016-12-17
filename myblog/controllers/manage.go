package controllers

import (
	"github.com/astaxie/beego"
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
		topics, err := tools.GetAllTopics()
		if err != nil {
			beego.Error(err)
		}
		c.Data["Topics"] = topics
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
		c.TplName = "manage_add.html"
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
	oldtitle := c.Input().Get("oldtitle")
	title := c.Input().Get("title")
	content := c.Input().Get("content")
	category := c.Input().Get("category")
	createdtime := c.Input().Get("createdtime")

	//下面是将其保存下来
	if createdtime == "" {
		//创建文件夹，这里是创建分类
		tools.Createdir("topic", category)
		//当前时间
		nowtime := time.Now().Format("2006-01-02")
		//写入内容到md
		mdname := nowtime + "#" + title + ".md"
		tools.Writefile(category, mdname, content)
	} else {
		//创建文件夹，这里是创建分类
		tools.Createdir("topic", category)
		//写入内容到md
		mdname := createdtime + "#" + title + ".md"
		tools.Writefile(category, mdname, content)
		//顺便删除老的md
		err := tools.DeleteTopic(oldtitle, createdtime, category)
		if err != nil {
			beego.Error(err)
		}
	}

	c.Redirect("/manage", 302)
}

// 修改文章
func (c *ManageController) Modify() {

	if !checkAccount(c.Ctx) {
		c.Redirect("/", 302)
		return
	}

	c.TplName = "manage_modify.html"
	title := c.Input().Get("title")

	//查看文章详细
	var err error
	topic, err := tools.GetTopicsContent(title)
	if err != nil {
		beego.Error(err)
		c.Redirect("/", 302)
		return
	}

	c.Data["Topic"] = topic[0]
}

// 删除文章
func (c *ManageController) Delete() {
	if !checkAccount(c.Ctx) {
		c.Redirect("/", 302)
		return
	}

	title := c.Input().Get("title")
	category := c.Input().Get("category")
	createdtime := c.Input().Get("createdtime")

	err := tools.DeleteTopic(title, createdtime, category)
	if err != nil {
		beego.Error(err)
	}
	c.Redirect("/manage", 302)
	return
}
