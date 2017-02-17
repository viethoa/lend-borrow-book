package book

import (
  "fmt"
  "log"
  "net/http"
  "strconv"
  "lend-borrow-book/models"
  "lend-borrow-book/repository"
)

func Borrow(w http.ResponseWriter, req *http.Request) {
  err := req.ParseForm()
  if err != nil {
    log.Fatalf("ParseForm error: %s", err.Error())
    return
  }

  // Check params
  email := req.Header.Get("email")
  if (len(email) == 0) {
    w.WriteHeader(400)
    w.Write([]byte("Missing email on header"))
    return
  }
  bookIDStr := req.FormValue("book_id");
  bookID, err := strconv.Atoi(bookIDStr);
  if (err != nil) {
    w.WriteHeader(422)
    w.Write([]byte("Invalid slot ID"))
    return
  }

  // Check book has exists
  bookRepo := repository.Book{}
  book, err := bookRepo.Get(bookID)
  if (err != nil) {
    w.WriteHeader(422)
    w.Write([]byte("This book not exists"))
    return;
  }

  Transaction := models.Transaction {
    BookName: book.Name,
    Email: email,
  }

  TransactionRepo := repository.Transaction{}
  if err = TransactionRepo.Create(&Transaction); err != nil {
    w.WriteHeader(500)
    fmt.Fprintf(w, "Cannot persist. Error %s", err.Error())
    return
  }

  fmt.Fprintf(w, "Email %s have borrowing book %s", email, book.Name)
}
