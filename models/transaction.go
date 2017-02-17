package models

import "time"

type Transaction struct {
  ID                int         `gorm:"id"`
  BookName          string      `gorm:"book_name"`
  Email             string      `gorm:"email"`
  CreatedAt         time.Time   `gorm:"created_at"`
}