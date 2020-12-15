package sysinit

import (
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB
var err error

func InitGorm() {
	DB, err = gorm.Open("mysql", "root:123456@(192.168.1.145:3306)/bee?charset=utf8")
	if err != nil {
		beego.Error(err)
		return
	}
}
