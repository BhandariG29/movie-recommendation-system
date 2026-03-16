package models

import "time"

type WatchHistory struct{
	W_ID				uint		`json:"w_id" gorm:"primaryKey"`
	U_ID				uint		`json:"u_id"`
	M_ID				uint		`json:"m_id"`
	Rating				float32		`json:"rating"`
	WatchPercentage		float32		`json:"watch_percentage"`
	WatchedAt			time.Time	`json:"watched_at"`
}