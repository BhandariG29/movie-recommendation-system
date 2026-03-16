package models

type Genre struct {
	G_ID   	uint   `json:"g_id" gorm:"primaryKey"`
	Name 	string `json:"name"`
}