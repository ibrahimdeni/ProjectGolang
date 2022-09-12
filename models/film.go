package models

type Film struct {
	ID            int                  `json:"id" gorm:"primary_key:auto_increment"`
	Title         string               `json:"title" gorm:"type: varchar(255)"`
	Thumbnailfilm string               `json:"thumbnailfilm" gorm:"type:varchar(255)"`
	Year          string               `json:"year" gorm:"type: varchar(255)" `
	Description   string               `json:"description" gorm:"type: varchar(255)"`
	User          UsersProfileResponse `json:"user"`
	UserID        int                  `json:"-"`
}

type FilmResponse struct {
	ID            int                  `json:"id" gorm:"primary_key:auto_increment"`
	Title         string               `json:"title" gorm:"type: varchar(255)"`
	Thumbnailfilm string               `json:"thumbnailfilm" gorm:"type:varchar(255)"`
	Year          string               `json:"year" gorm:"type: varchar(255)" `
	Description   string               `json:"description" gorm:"type: varchar(255)"`
	User          UsersProfileResponse `json:"user"`
	UserID        int                  `json:"-"`
}

func (FilmResponse) TableName() string {
	return "films"
}
