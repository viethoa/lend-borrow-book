package book

import (
  "fmt"
  "log"
  "net/http"
  "lend-borrow-book/repository"
  "encoding/json"
)

func GetAll(w http.ResponseWriter, req *http.Request) {
  err := req.ParseForm()
  if err != nil {
    log.Fatalf("ParseForm error: %s", err.Error())
    return
  }

  repo := repository.Book{}
  books, err := repo.GetAll()
  if err != nil {
    w.WriteHeader(500)
    fmt.Fprintf(w, "Cannot persist. Error %s", err.Error())
    return
  }

  data, _ := json.Marshal(books)
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(200)
  w.Write(data)

  fmt.Printf("Get all books: %v", books)
}
