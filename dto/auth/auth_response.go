package authdto

type LoginResponse struct {
	ID       int    `json:"id"`
	Name     string `gorm:"type: varchar(255)" json:"name"`
	Password string `gorm:"type: varchar(255)" json:"password" validate:"required"`
	Email    string `gorm:"type: varchar(255)" json:"email"`
	Token    string `gorm:"type: varchar(255)" json:"token"`
}