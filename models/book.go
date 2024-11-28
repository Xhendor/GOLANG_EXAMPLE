package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json:"title" validate:"required,min=1,max=200"`
	Author string `json:"author" validate:"required,min=1,max=100"`
	ISBN   string `json:"isbn" validate:"required,min=10,max=13"`
}

// Validate returns error if book data is invalid
func (b *Book) Validate() error {
	if b.Title == "" {
		return ErrEmptyTitle
	}
	if b.Author == "" {
		return ErrEmptyAuthor
	}
	if b.ISBN == "" {
		return ErrEmptyISBN
	}
	return nil
}
