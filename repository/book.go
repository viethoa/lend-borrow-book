package repository

import (
  "lend-borrow-book/db"
  "lend-borrow-book/models"
)

// Book represents repository for corresponding Book model
type Book struct{}

// Create persists Request model to database
func (Book) Create(book *models.Book) error {
  return db.DB.Save(book).Error
}

func (Book) GetAll() ([]models.Book, error) {
  books := []models.Book{}
  err := db.DB.Find(&books).Error
  return books, err
}

func (Book) Get(id int) (models.Book, error) {
  book := models.Book{}
  err := db.DB.First(&book, 10).Error
  return book, err
}
