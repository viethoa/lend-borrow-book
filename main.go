package main

import (
	"net/http"
  "fmt"
	"log"
	"lend-borrow-book/routers"
  "lend-borrow-book/env"
)

func main() {
	r := routers.New()
  port := env.Get("PORT")
  fmt.Printf("Serving on host: localhost: %s", port)
	log.Fatal(http.ListenAndServe(":" + port, r))
}

// CREATE TABLE books (
//    id BIGSERIAL PRIMARY KEY NOT NULL,
//    name VARCHAR(100) NOT NULL,
//    short_description VARCHAR(250),
//    category VARCHAR(100),
//    status VARCHAR(100),
//    suggestion VARCHAR(250),
//    created_at TIMESTAMP
// );