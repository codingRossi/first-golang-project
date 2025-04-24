package models

type User struct {
	id   int    `gorm:"type:int;primary_key"`
	name string `gorm:"type:varchar(255)"`
	email string `gorm:"type:varchar(255)"`
}