package services

import (
	"bee2dbtask/db"
	"bee2dbtask/models"
	// "log"
	// "reflect"
)

func GettingData(tab []map[string]interface{}) {
	db.ConnectToDB()
	db.AutoMigrate()
	details := []models.Book{}
	for _, eachRecord := range tab {
		//log.Println(eachRecord["Id"], reflect.TypeOf(eachRecord["Id"]))
		tabA := models.Book{}
		p := int(eachRecord["Id"].(float64))
		//log.Println(reflect.TypeOf(p))
		tabA.Id = p
		tabA.Title = eachRecord["Title"].(string)
		tabA.Author = eachRecord["Author"].(string)
		q := int(eachRecord["Year"].(float64))
		tabA.Year = q
		tabA.Price = eachRecord["Price"].(float64)
		details = append(details, tabA)
	}
	db.Connection.Save(&details)
}
