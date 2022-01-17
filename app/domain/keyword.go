package domain

import "time"

type Keyword struct {
	ID           uint      `gorm:"primary_key" json:"id"`
	Adword       string    `json:"adword"`
	Link         string    `json:"link"`
	SearchResult string    `json:"search_result"`
	Status       string    `json:"status"`
	UserId       uint      `json:"user_id"`
	User         User      `gorm:"foreignKey:UserId;references:ID"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
