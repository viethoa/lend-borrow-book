package controller

import "net/http"

// Root is the entrypoint of the app
func Root(w http.ResponseWriter, req *http.Request) {
  w.Write([]byte("pong"))
}