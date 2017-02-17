package repository_test

// Create persists Request model to database
func TestCreate(book *models.Book) error {
  return db.DB.Save(book).Error
}
