package models


type Author struct{
	Id string `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Bio string `json:"bio"`
}