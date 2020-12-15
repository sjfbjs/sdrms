package controllers

import (
	"github.com/astaxie/beego"
	"github.com/sjfbjs/sdrms/models"
	"github.com/sjfbjs/sdrms/sysinit"
	"time"
)

type MySqlController struct {
	//BaseController
	beego.Controller
}

//查询
func (c *MySqlController) Get() {
	user := []models.User{}

	// user := []sysinit.User{}
	sysinit.DB.Find(&user)
	c.Data["json"] = user
	c.ServeJSON()
	c.TplName = "mysql.html"
}

//新增
func (c *MySqlController) Add() {
	user := models.User{
		Name:    "张三",
		Age:     22,
		AddTime: int(time.Now().Unix()),
	}
	sysinit.DB.Create(&user)
	c.Ctx.WriteString("增加数据成功")
}

//更新
func (c *MySqlController) Update() {
	user := models.User{ID: 4}
	sysinit.DB.First(&user)
	user.Name = "周华建"
	sysinit.DB.Save(&user)
	c.Ctx.WriteString("更新数据成功")

}

//删除
func (c *MySqlController) Delete() {
	user := models.User{ID: 5}
	sysinit.DB.Delete(&user)
	c.Ctx.WriteString("删除数据库成功")

}
