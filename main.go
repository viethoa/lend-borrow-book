package main

import (
	"net/http"
	"log"
	"lend-borrow-book/routers"
  "lend-borrow-book/env"

  "database/sql"
  _ "github.com/lib/pq"
)

func main() {
  db, err := sql.Open("postgres", "user=VietHoa password=viethoa dbname=lendborrowbook sslmode=disable")
  if err != nil {
    log.Fatal(err)
  }

  var bookID int
  error := db.QueryRow(`INSERT INTO book (name, short_description, category, status, suggestion) VALUES ('Test Book', 'test desciption', 'none', 'avaiable', 'dont read it') RETURNING id`).Scan(&bookID)
  if error != nil {
    log.Fatal(error)
  }

	r := routers.New()
  port := env.Get("PORT")
	log.Fatal(http.ListenAndServe(":" + port, r))
}