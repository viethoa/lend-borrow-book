package main

import (
	"net/http"
	"log"
	"lend-borrow-book/routers"
)

func main() {
	r := routers.New()
	log.Fatal(http.ListenAndServe(":8080", r))
}