package controllers

import (
	"github.com/astaxie/beego"
	"myblog/tools"
	"strconv"
	"time"
)

type TopicController struct {
	beego.Controller
}

//分页
var startpage = 0

func (c *TopicController) Get() {

	//登陆状态就可以跳转
	c.TplName = "topic.html"
	//检查是不是登陆状态
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	topics, err := tools.GetAllTopicsTopic()
	if err != nil {
		beego.Error(err)
	}

	//每页显示条数
	Display := 5
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

}

//查看文章详情
func (c *TopicController) View() {

	//检查是不是登陆状态
	c.Data["IsLogin"] = checkAccount(c.Ctx)

	// 获取title
	title := c.Input().Get("title")

	//查看文章详细
	topic, err := tools.GetTopicsContent(title)
	if err != nil {
		beego.Error(err)
	}
	c.Data["Topic"] = topic[0]
	//第一次可能没有评论
	c.Data["Content"], err = tools.ReadComment(title)
	if err != nil {
		beego.Error(err)
		c.Data["Content"] = ""
	}
	//查看文章详细
	c.TplName = "topic_view.html"
}

func (c *TopicController) Add() {
	//获取表单
	title := c.Input().Get("title")
	email := c.Input().Get("email")
	content := c.Input().Get("content")
	//当前时间
	nowtime := time.Now().Format("2006-01-02")

	if email == "420439007@qq.tyb" {
		email = "作者TTyb"
	}
	err := tools.WriteComment(title, email, nowtime, content)
	if err != nil {
		beego.Error(err)
	}

	c.Data["Content"], err = tools.ReadComment(title)
	if err != nil {
		beego.Error(err)
	}
	c.Redirect("/topic/view?title=" + title, 302)
}