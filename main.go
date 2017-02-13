package main

import (
	"net/http"
	"log"
	"LandBorrowBook/routers"
)

func main() {
	r := routers.New()
	log.Fatal(http.ListenAndServe(":8080", r))
}