package models

import (
	"time"
)

// Book represents a book entity
type Book struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at,omitempty" gorm:"index"`
	Title     string    `json:"title" validate:"required,min=1,max=200" example:"The Go Programming Language"`
	Author    string    `json:"author" validate:"required,min=1,max=100" example:"Alan A. A. Donovan"`
	ISBN      string    `json:"isbn" validate:"required,min=10,max=13" example:"978-0134190440"`
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
