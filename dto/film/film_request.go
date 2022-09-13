package film

type FilmRequest struct {
	ID            int    `json:"id" gorm:"primary_key:auto_increment"`
	Title         string `json:"title" from:"title"  gorm:"type: varchar(255)"`
	Thumbnailfilm string `json:"thumbnailfilm" gorm:"type: varchar(255)"`
	Year          string `json:"year" form:"year" gorm:"type: text" validate:"required" `
	Description   string `json:"description" gorm:"type: varchar(255)"`
	CategoryID    int    `json:"category_id"`
}
