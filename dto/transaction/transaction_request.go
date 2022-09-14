package transaction

import "time"

type TransactionRequest struct {
	ID       int       `json:"id" gorm:"primary_key:auto_increment"`
	Stardate time.Time `json:"created_at"`
	Duedate  time.Time `json:"updated_at"`
	UserID   int       `json:"user_id" gorm:"type: int"`
	Attache   string    `json:"attach" gorm:"type: text"`
	Status   string    `json:"status" gorm:"type: text"`
}
