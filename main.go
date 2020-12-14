package main

import (
	_ "github.com/sjfbjs/sdrms/routers"
	_ "github.com/sjfbjs/sdrms/sysinit"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
