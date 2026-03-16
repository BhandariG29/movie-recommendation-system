package models

type Actor struct {
	A_ID   	uint   `json:"a_id" gorm:"primaryKey"`
	Name 	string `json:"name"`
}