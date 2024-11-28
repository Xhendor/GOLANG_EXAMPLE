package models

import "errors"

var (
	ErrEmptyTitle  = errors.New("title cannot be empty")
	ErrEmptyAuthor = errors.New("author cannot be empty")
	ErrEmptyISBN   = errors.New("ISBN cannot be empty")
	ErrBookNotFound = errors.New("book not found")
)
