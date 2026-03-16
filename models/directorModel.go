package models

type Director struct {
	D_ID   	uint   `json:"d_id" gorm:"primaryKey"`
	Name 	string `json:"name"`
}