package main

import (
	"net/http"
	"log"
	"lend-borrow-book/routers"
  "lend-borrow-book/env"
)

func main() {
	r := routers.New()
  port := env.Get("PORT")
	log.Fatal(http.ListenAndServe(":" + port, r))
}