package routers

import (
	"github.com/astaxie/beego"
	"github.com/sjfbjs/sdrms/controllers"
)

func init() {
	//测试路由?
	beego.Router("/course/myindex", &controllers.TestController{}, "*:Index")
	////基础路由，可用namespace整改
	ns := beego.NewNamespace("/sjf",
		beego.NSRouter("/index", &controllers.TestController{}, "*:Index"),
	)
	beego.AddNamespace(ns)
	v1ns := beego.NewNamespace("v1",
		beego.NSRouter("/index", &controllers.TestController{}, "*:Index"),
	)
	beego.AddNamespace(v1ns)

	//课程路由
	beego.Router("/course/index", &controllers.CourseController{}, "*:Index")
	beego.Router("/course/add", &controllers.CourseController{}, "*:Add")
	beego.Router("/course/datagrid", &controllers.CourseController{}, "Get,Post:DataGrid")
	beego.Router("/course/edit/?:id", &controllers.CourseController{}, "Get,Post:Edit")
	beego.Router("/course/delete", &controllers.CourseController{}, "Post:Delete")
	beego.Router("/course/updateseq", &controllers.CourseController{}, "Post:UpdateSeq")
	beego.Router("/course/uploadimage", &controllers.CourseController{}, "Post:UploadImage")

	//用户角色路由
	beego.Router("/role/index", &controllers.RoleController{}, "*:Index")
	beego.Router("/role/datagrid", &controllers.RoleController{}, "Get,Post:DataGrid")
	beego.Router("/role/edit/?:id", &controllers.RoleController{}, "Get,Post:Edit")
	beego.Router("/role/delete", &controllers.RoleController{}, "Post:Delete")
	beego.Router("/role/datalist", &controllers.RoleController{}, "Post:DataList")
	beego.Router("/role/allocate", &controllers.RoleController{}, "Post:Allocate")
	beego.Router("/role/updateseq", &controllers.RoleController{}, "Post:UpdateSeq")

	//资源路由
	beego.Router("/resource/index", &controllers.ResourceController{}, "*:Index")
	beego.Router("/resource/treegrid", &controllers.ResourceController{}, "POST:TreeGrid")
	beego.Router("/resource/edit/?:id", &controllers.ResourceController{}, "Get,Post:Edit")
	beego.Router("/resource/parent", &controllers.ResourceController{}, "Post:ParentTreeGrid")
	beego.Router("/resource/delete", &controllers.ResourceController{}, "Post:Delete")
	//快速修改顺序
	beego.Router("/resource/updateseq", &controllers.ResourceController{}, "Post:UpdateSeq")

	//通用选择面板
	beego.Router("/resource/select", &controllers.ResourceController{}, "Get:Select")
	//用户有权管理的菜单列表（包括区域）
	beego.Router("/resource/usermenutree", &controllers.ResourceController{}, "POST:UserMenuTree")
	beego.Router("/resource/checkurlfor", &controllers.ResourceController{}, "POST:CheckUrlFor")

	//后台用户路由
	beego.Router("/backenduser/index", &controllers.BackendUserController{}, "*:Index")
	beego.Router("/backenduser/datagrid", &controllers.BackendUserController{}, "POST:DataGrid")
	beego.Router("/backenduser/edit/?:id", &controllers.BackendUserController{}, "Get,Post:Edit")
	beego.Router("/backenduser/delete", &controllers.BackendUserController{}, "Post:Delete")
	//后台用户中心
	beego.Router("/usercenter/profile", &controllers.UserCenterController{}, "Get:Profile")
	beego.Router("/usercenter/basicinfosave", &controllers.UserCenterController{}, "Post:BasicInfoSave")
	beego.Router("/usercenter/uploadimage", &controllers.UserCenterController{}, "Post:UploadImage")
	beego.Router("/usercenter/passwordsave", &controllers.UserCenterController{}, "Post:PasswordSave")

	beego.Router("/home/index", &controllers.HomeController{}, "*:Index")
	beego.Router("/home/login", &controllers.HomeController{}, "*:Login")
	beego.Router("/home/dologin", &controllers.HomeController{}, "Post:DoLogin")
	beego.Router("/home/logout", &controllers.HomeController{}, "*:Logout")
	beego.Router("/home/datareset", &controllers.HomeController{}, "Post:DataReset")

	beego.Router("/home/404", &controllers.HomeController{}, "*:Page404")
	beego.Router("/home/error/?:error", &controllers.HomeController{}, "*:Error")
	beego.Router("/", &controllers.HomeController{}, "*:Index")

	beego.Router("/mysql", &controllers.MySqlController{})
	beego.Router("/mysql/add", &controllers.MySqlController{}, "get:Add")
	beego.Router("/mysql/update", &controllers.MySqlController{}, "get:Update")
	beego.Router("/mysql/delete", &controllers.MySqlController{}, "get:Delete")

	//ns :=
	//	web.NewNamespace("/v1",
	//		web.NSCond(func(ctx *context.Context) bool {
	//			if ctx.Input.Domain() == "api.beego.me" {
	//				return true
	//			}
	//			return false
	//		}),
	//		web.NSBefore(auth),
	//		web.NSGet("/notallowed", func(ctx *context.Context) {
	//			ctx.Output.Body([]byte("notAllowed"))
	//		}),
	//		web.NSRouter("/version", &AdminController{}, "get:ShowAPIVersion"),
	//		web.NSRouter("/changepassword", &UserController{}),
	//		web.NSNamespace("/shop",
	//			web.NSBefore(sentry),
	//			web.NSGet("/:id", func(ctx *context.Context) {
	//				ctx.Output.Body([]byte("notAllowed"))
	//			}),
	//		),
	//		web.NSNamespace("/cms",
	//			web.NSInclude(
	//				&controllers.MainController{},
	//				&controllers.CMSController{},
	//				&controllers.BlockController{},
	//			),
	//		),
	//	)
	////注册 namespace
	//web.AddNamespace(ns)

}
