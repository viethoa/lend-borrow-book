package book

import (
  "fmt"
  "log"
  "net/http"
  "lend-borrow-book/models"
  "lend-borrow-book/repository"
)

func Create(w http.ResponseWriter, req *http.Request) {
  err := req.ParseForm()
  if err != nil {
    log.Fatalf("ParseForm error: %s", err.Error())
    return
  }

  name := req.FormValue("name")
  if (len(name) == 0) {
    w.WriteHeader(400)
    w.Write([]byte("Missing book name"))
    return
  }
  category := req.FormValue("category")
  if (len(category) == 0) {
    w.WriteHeader(400)
    w.Write([]byte("Missing category for this book"))
    return
  }

  bookModel := models.Book{
    Name: name,
    ShortDescription: req.FormValue("short_description"),
    Category: category,
    Suggestion: req.FormValue("suggestion"),
    Status: "avaiable",
  }

  repo := repository.Book{}
  if err = repo.Create(&bookModel); err != nil {
    w.WriteHeader(500)
    fmt.Fprintf(w, "Cannot persist. Error %s", err.Error())
    return
  }

  fmt.Fprintf(w, "Book: %s", name)
}