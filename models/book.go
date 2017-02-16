package models

import "time"

type Book struct {
  ID                int         `gorm:"id"`
  Name              string      `gorm:"name"`
  ShortDescription  string      `gorm:"short_description"`
  Category          string      `gorm:"category"`
  Status            string      `gorm:"status"`
  Suggestion        string      `gorm:"suggestion"`
  CreatedAt         time.Time   `gorm:"created_at"`
}
