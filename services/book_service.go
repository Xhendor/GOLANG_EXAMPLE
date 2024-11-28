package services

import (
	"go-crud/models"
	"gorm.io/gorm"
)

type BookService struct {
	db *gorm.DB
}

func NewBookService(db *gorm.DB) *BookService {
	return &BookService{db: db}
}

func (s *BookService) GetAll() ([]models.Book, error) {
	var books []models.Book
	result := s.db.Find(&books)
	return books, result.Error
}

func (s *BookService) GetByID(id uint) (*models.Book, error) {
	var book models.Book
	result := s.db.First(&book, id)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, models.ErrBookNotFound
	}
	return &book, result.Error
}

func (s *BookService) Create(book *models.Book) error {
	return s.db.Create(book).Error
}

func (s *BookService) Update(book *models.Book) error {
	result := s.db.Model(book).Updates(book)
	if result.RowsAffected == 0 {
		return models.ErrBookNotFound
	}
	return result.Error
}

func (s *BookService) Delete(id uint) error {
	result := s.db.Delete(&models.Book{}, id)
	if result.RowsAffected == 0 {
		return models.ErrBookNotFound
	}
	return result.Error
}
