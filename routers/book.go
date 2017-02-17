package routers

import (
  bookapi "lend-borrow-book/controllers/api/book"

  "github.com/pressly/chi"
)

func bookAPI(r chi.Router) {
  r.Post("/create", bookapi.Create)
  r.Get("/getAll", bookapi.GetAll)
  r.Get("/borrow", bookapi.Borrow)
}
