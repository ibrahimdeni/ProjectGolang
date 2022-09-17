package models

// import "time"

type Transaction struct {
	ID       int       `json:"id" gorm:"primary_key:auto_increment"`
	Stardate string `json:"startdate"`
	Duedate  string `json:"duedate"`
	Attache  string    `json:"attach" gorm:"type: text"`
	Status   string    `json:"status" gorm:"type: text"`
	User	 UsersProfileResponse	`json:"user"`
	UserID	 int		`json:"-"`
}
