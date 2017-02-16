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

