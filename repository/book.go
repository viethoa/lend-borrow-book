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
