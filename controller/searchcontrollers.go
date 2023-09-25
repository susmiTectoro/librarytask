package controller

import (
	"bee2dbtask/db"
	"fmt"

	"github.com/beego/beego/v2/server/web"
)

type SearchController struct {
	web.Controller
}

func (s *SearchController) SearchLib() {
	data := s.Ctx.Request.Form
	//result := &models.Book{}
	abc := make([]map[string]interface{}, 0)
	for k, v := range data {
		s := fmt.Sprintf("%s:%s", k, v[0])
		fmt.Println(s)
		db.Connection.Raw("SELECT x.* FROM `user`.books x WHERE ?", s).Scan(&abc)
		fmt.Println(abc)
	}
	//fmt.Println(abc)
	s.Data["json"] = "success"
	s.ServeJSON()
}

// for _, val := range abc {
// 	switch x {
// 	case "Year":
// 		p, _ := strconv.Atoi(y)
// 		if int(val["year"].(int64)) == p {
// 			xy = append(xy, val)
// 		}
// 	case "Title":
// 		if val["title"] == y {
// 			xy = append(xy, val)
// 		}
// 	case "Author":
// 		if val["author"] == y {
// 			xy = append(xy, val)
// 		}
// 	case "MaxPrice":
// 		s, _ := strconv.ParseFloat(y, 64)
// 		p := val["price"].(float64)
// 		if p < s {
// 			xy = append(xy, val)
// 		}
// 	case "MinPrice":
// 		s, _ := strconv.ParseFloat(y, 64)
// 		p := val["price"].(float64)
// 		if p > s {
// 			xy = append(xy, val)
// 		}
// 	}
// }
// s.Data["json"] = xy
// s.ServeJSON()
