package main

import (
	//"bee2dbtask/db"
	_ "bee2dbtask/routers"

	"github.com/beego/beego/v2/server/web"
)

func main() {

	web.Run()

}
