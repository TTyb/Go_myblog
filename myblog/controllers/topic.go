package controllers

import (
	"github.com/astaxie/beego"
	"myblog/tools"
)

type TopicController struct {
	beego.Controller
}

func (c *TopicController) Get() {

	//登陆状态就可以跳转
	c.TplName = "topic.html"
	//检查是不是登陆状态
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	topics, err := tools.GetAllTopicsTopic()
	if err != nil {
		beego.Error(err)
	}
	c.Data["Topics"] = topics
}

//查看文章详情
func (c *TopicController) View() {

	//检查是不是登陆状态
	c.Data["IsLogin"] = checkAccount(c.Ctx)

	// 获取id
	title := c.Input().Get("title")

	//查看文章详细
	var err error
	topic, err := tools.GetTopicsContent(title)
	if err != nil {
		beego.Error(err)
	}
	c.Data["Topic"] = topic[0]
	//查看文章详细
	c.TplName = "topic_view.html"

}