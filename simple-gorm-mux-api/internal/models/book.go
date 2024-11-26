package models


type Book struct{
	Id string `json:"id" gorm:"primaryKey"`
	Title string `json:"title"`
	Genre string `json:"genre"`
}