package db

import (
	"bee2dbtask/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Connection *gorm.DB

func ConnectControlDB(host string, port int64, database string, user string, pass string) (dbConn *gorm.DB, err error) {
	dbConnString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, database)
	dbConn, _ = gorm.Open(mysql.Open(dbConnString))
	Connection = dbConn
	return
}
func ConnectToDB() *gorm.DB {
	conn, _ := ConnectControlDB("localhost", 3306, "user", "susmiTectoro", "susmiTectoro")
	fmt.Println("Connected to Control DB")
	Connection = conn
	return conn
}

func AutoMigrate() {
	var tables []interface{}
	tables = append(tables, &models.Book{})
	err := Connection.AutoMigrate(tables...)
	if err != nil {
		fmt.Println("Error occurred:", err)
	}
}
