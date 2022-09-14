package models

import "time"

type Transaction struct {
	ID       int       `json:"id" gorm:"primary_key:auto_increment"`
	Stardate time.Time `json:"created_at"`
	Duedate  time.Time `json:"updated_at"`
	Attache  string    `json:"attach" gorm:"type: text"`
	Status   string    `json:"status" gorm:"type: text"`
	User	 UsersProfileResponse	`json:"user"`
	UserID	 int		`json:"-"`
}
