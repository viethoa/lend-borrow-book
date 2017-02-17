package repository

import (
  "lend-borrow-book/db"
  "lend-borrow-book/models"
)

// LendBorrow represents repository for corresponding LendBorrow model
type Transaction struct{}

// Create persists Request model to database
func (Transaction) Create(transaction *models.Transaction) error {
  return db.DB.Save(transaction).Error
}
