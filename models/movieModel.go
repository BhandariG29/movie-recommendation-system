package models



type Movie struct{
	M_ID			uint		`json:"m_id" gorm:"primaryKey"`
	Title			string		`json:"title"`
	ReleaseYear		int			`json:"release_year"`
	D_ID			uint		`json:"d_id"`
}