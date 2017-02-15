package main

import (
	"net/http"
	"log"
	"lend-borrow-book/routers"

  "os"
)

func main() {
	r := routers.New()

  port := os.Getevn("PORT")
  if (port == "") {
    port = "8080"
  }

	log.Fatal(http.ListenAndServe(":" + port, r))
}