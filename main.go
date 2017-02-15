package main

import (
	"net/http"
	"log"
	"LendBorrowBook/routers"
)

func main() {
	r := routers.New()
	log.Fatal(http.ListenAndServe(":8080", r))
}