package sysinit

import (
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB
var err error

func InitGorm() {
	//数据库类别
	dbType := beego.AppConfig.String("db_type")
	//连接名称
	//dbAlias := beego.AppConfig.String(dbType + "::db_alias")
	//数据库名称
	dbName := beego.AppConfig.String(dbType + "::db_name")
	//数据库连接用户名
	dbUser := beego.AppConfig.String(dbType + "::db_user")
	//数据库连接用户名
	dbPwd := beego.AppConfig.String(dbType + "::db_pwd")
	//数据库IP（域名）
	dbHost := beego.AppConfig.String(dbType + "::db_host")
	//数据库端口
	dbPort := beego.AppConfig.String(dbType + "::db_port")
	dbCharSet := beego.AppConfig.String(dbType + "::db_charset")
	DB, err = gorm.Open(dbType, dbUser+":"+dbPwd+"@("+dbHost+":"+dbPort+")/"+dbName+"?charset="+dbCharSet)
	if err != nil {
		beego.Error(err)
		return
	}
}
