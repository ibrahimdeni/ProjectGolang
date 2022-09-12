package models

type Category struct {
	ID     int          `json:"id"`
	Name   string       `json:"name" form:"name" validate:"required"`
	FilmID int          `json:"-"`
	Film   FilmResponse `json:"film"`
}
