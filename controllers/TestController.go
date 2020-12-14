package controllers

import "github.com/astaxie/beego"

type TestController struct {
	beego.Controller
}

func (c TestController) Index() {
	c.Ctx.Output.Body([]byte("testindex"))
}
