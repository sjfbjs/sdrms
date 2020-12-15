package main

import (
	_ "github.com/sjfbjs/sdrms/routers"
	"github.com/sjfbjs/sdrms/sysinit"
	_ "github.com/sjfbjs/sdrms/sysinit"

	"github.com/astaxie/beego"
)

func main() {
	//beego.SetLogger("file",`{"filename":"logs/test.log"}`)
	//beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
	//关闭数据库连接
	defer sysinit.DB.Close()
}
