package transaction

import (
	"dumbflix/models"
	// "time"
)

type TransactionResponse struct {
	ID       	int       	`json:"id" gorm:"primary_key:auto_increment"`
	Stardate 	string 	`json:"startdate"`
	Duedate  	string 	`json:"duedate"`
	User		models.UsersProfileResponse	`json:"userId"`
	// UserID   	int       	`json:"user_id" gorm:"type: int"`
	Attache   	string   	`json:"attach" gorm:"type: text"`
	Status   	string    	`json:"status" gorm:"type: text"`
}
