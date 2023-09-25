package controller

import (
	"bee2dbtask/db"
	"bee2dbtask/services"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/beego/beego/v2/server/web"
)

type LibraryController struct {
	web.Controller
}

func (u *LibraryController) CreateLib() {
	records := getdata(u.Ctx.Input.RequestBody)

	services.GettingData(records)
	u.Data["json"] = records
	u.ServeJSON()
}

func (u *LibraryController) GetLib() {
	abc := make([]map[string]interface{}, 0)
	db.Connection.Raw("SELECT x.* FROM `user`.books x").Scan(&abc)
	u.Data["json"] = abc
	u.ServeJSON()
}

func (u *LibraryController) GetLibData() {
	abc := make([]map[string]interface{}, 0)
	db.Connection.Debug().Raw("SELECT x.* FROM `user`.books x").Scan(&abc)
	fmt.Println(abc)

	inLine := u.Ctx.Input.Param(":id")
	s, _ := strconv.Atoi(inLine)
	for _, value := range abc {
		if int(value["id"].(int64)) == s {
			u.Data["json"] = value
			u.ServeJSON()
		}
	}
}

func (u *LibraryController) PutLibData() {
	abc := getdata(u.Ctx.Input.RequestBody)
	services.GettingData(abc)
	u.Data["json"] = abc
	u.ServeJSON()
}

func (u *LibraryController) DelLibData() {
	inLine := u.Ctx.Input.Param(":id")
	str := strings.Split(inLine, ",")
	for _, v := range str {
		x, _ := strconv.Atoi(v)
		abc := make([]map[string]interface{}, 0)
		db.Connection.Raw("DELETE  FROM `user`.books  WHERE id= ?", x).Scan(&abc)
	}
	u.Data["json"] = "Book Deleted"
	u.ServeJSON()
}
func getdata(bytedata []byte) []map[string]interface{} {
	var data []map[string]interface{}
	json.Unmarshal(bytedata, &data)
	return data
}
