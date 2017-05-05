package controllers

import (
	"github.com/astaxie/beego"
	"time"
	"myblog/tools"
	"strconv"
)

type ManageController struct {
	beego.Controller
}

//跳转到管理页面
func (c *ManageController) Get() {
	if checkAccount(c.Ctx) == true {
		c.Data["IsLogin"] = true
		c.TplName = "manage_topic.html"
		topics, err := tools.GetAllTopics()
		if err != nil {
			beego.Error(err)
		}
		//每页显示条数
		Display := 10
		c.Data["Display"] = Display

		//一共都有多少条数据
		Count := len(topics)
		c.Data["Count"] = Count

		//首页
		FirstPage := 0
		c.Data["FirstPage"] = FirstPage + 1

		//末页
		LastPage := Count / Display
		if Count % Display != 0 {
			LastPage = LastPage + 1
			c.Data["LastPage"] = LastPage
		} else {
			c.Data["LastPage"] = LastPage
		}

		//接受点击回来的数据
		page := c.Input().Get("page")
		// 首页
		if page == strconv.Itoa(FirstPage) {
			startpage = 0
			c.Data["Topics"] = topics[startpage:startpage + Display]
		} else if page == strconv.Itoa(LastPage) {
			//末页
			temp := Count % Display
			startpage = Count - temp
			c.Data["Topics"] = topics[startpage:]
		} else if page != "" {
			if page != "0" {
				temp, err := strconv.Atoi(page)
				startpage = (temp - 1) * Display
				c.Data["Topics"] = topics[startpage:startpage + Display]
				if err != nil {
					beego.Error(err)
				}
			}
		} else {
			//什么都不是那么就显示首页
			startpage = 0
			var err error
			c.Data["Topics"] = topics[startpage:startpage + Display]
			if err != nil {
				beego.Error(err)
			}
		}

		//跳转页
		skippage := c.Input().Get("skippage")
		if skippage != "" {
			skip, err := strconv.Atoi(skippage)
			if err != nil {
				beego.Error(err)
				startpage = 0
				c.Data["Topics"] = topics[startpage:startpage + Display]
			} else {
				startpage = (skip - 1) * Display
				if startpage + Display >= Count {
					tempnum := Count - Count % Display
					c.Data["Topics"] = topics[tempnum:Count]

				} else {
					c.Data["Topics"] = topics[startpage:startpage + Display]
				}
			}
		}

		//显示当前为第几页
		Nowpage := (startpage / Display) + 1
		if startpage % Display != 0 {
			Nowpage = Nowpage + 1
		}
		c.Data["Nowpage"] = Nowpage

		//上一页
		PrevPage := Nowpage - 1
		if Nowpage <= 1 {
			PrevPage = 1
			c.Data["PrevPage"] = PrevPage
		} else {
			c.Data["PrevPage"] = PrevPage
		}

		//下一页
		NextPage := Nowpage + 1
		if Nowpage >= LastPage {
			NextPage = LastPage
			c.Data["NextPage"] = NextPage
		} else {
			c.Data["NextPage"] = NextPage
		}

		//c.Data["Topics"] = topics
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
		c.TplName = "manage_addtopic.html"
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
	if createdtime == ""{
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

	c.Redirect("/manage/topic", 302)
}

// 修改文章
func (c *ManageController) Modify() {

	if !checkAccount(c.Ctx) {
		c.Redirect("/", 302)
		return
	}

	c.TplName = "manage_modifytopic.html"
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
	c.Redirect("/manage/topic", 302)
	return
}

// 管理日志
func (c *ManageController) Log() {

	//检查是不是登陆状态
	if checkAccount(c.Ctx) == true {
		c.Data["IsLogin"] = true
		//登陆状态就可以跳转
		c.TplName = "manage_log.html"
		c.Data["Content"] = tools.GetLogContent()
	} else {
		//不是登陆状态就老老实实的去登陆
		c.TplName = "home.html"
	}

}
