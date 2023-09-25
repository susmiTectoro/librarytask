package models

type Book struct {
	Id     int    `gorm:"Primarykey"`
	Title  string `gorm:"type:varchar(20)"`
	Author string `gorm:"type:varchar(20)"`
	Year   int
	Price  float64
}
