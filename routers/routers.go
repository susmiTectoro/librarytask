package routers

import (
	//"bee2dbtask/controller"

	"bee2dbtask/controller"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	web.Router("/User/info", &controller.LibraryController{}, "post:CreateLib")
	web.Router("/User/getinfo", &controller.LibraryController{}, "Get:GetLib")
	web.Router("/User/getinfo/:id", &controller.LibraryController{}, "Get:GetLibData")
	web.Router("/User/putinfo/:id", &controller.LibraryController{}, "Put:PutLibData")
	web.Router("/User/delinfo/:id", &controller.LibraryController{}, "Delete:DelLibData")
	web.Router("/User/getInfo/search/value", &controller.SearchController{}, "post:SearchLib")
}
