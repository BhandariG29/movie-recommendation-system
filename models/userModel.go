package models

import "time"

type User struct{
	U_ID		uint		`json:"u_id" gorm:"primaryKey"`
	Name		string		`json:"name"`
	Email		string		`json:"email"`
	CreatedAt	time.Time	`json:"created_at"`
}
