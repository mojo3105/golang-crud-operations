package models

type Book struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Title  string `json:"title" validate:"required"`
	Author string `json:"author" validate:"required"`
}