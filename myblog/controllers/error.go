package controllers

import (
    "github.com/astaxie/beego"
)

type ErrorController struct {
    beego.Controller
}

func (c *ErrorController) Error404() {
    c.TplName = "error.tpl"
}

//其他错误
//func (c *ErrorController) Error501() {
//    c.Data["content"] = "server error"
//    c.TplName = "501.tpl"
//}
//
//
//func (c *ErrorController) ErrorDb() {
//    c.Data["content"] = "database is now down"
//    c.TplName = "dberror.tpl"
//}