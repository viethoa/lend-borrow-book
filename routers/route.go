package routers

import (
  "lend-borrow-book/controllers"

  "github.com/pressly/chi"
  "github.com/pressly/chi/middleware"
)

var r chi.Router = chi.NewRouter()

func init() {
  r.Use(middleware.RealIP)
  r.Use(middleware.Logger)
  r.Use(middleware.Recoverer)
}

// New returns unexported chi Router instance with all mapped route
func New() chi.Router {
  r.Route("/", html)
  r.Route("/api", api)
  return r
}

func html(r chi.Router) {
  r.Get("/", controller.Root)
}

func api(r chi.Router) {
  r.Route("/book", bookAPI)
}
