package controllers

type TestController struct {
	BaseController
}

func (c TestController) Index() {
	//c.Ctx.Output.Body([]byte("testindex"))
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.controllerName + "." + c.actionName)
	c.setTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "course/index_headcssjs.html"
	c.LayoutSections["footerjs"] = "course/index_footerjs.html"
	//页面里按钮权限控制
	c.Data["canEdit"] = c.checkActionAuthor("TestController", "Edit")
	c.Data["canDelete"] = c.checkActionAuthor("TestController", "Delete")

}
