package api

import (
  "encoding/json"
  "net/http"
)

type errMsg struct {
  Msg string `json:"msg"`
}

type apiError struct {
  Error errMsg `json:"error"`
}

// Error is similar to http.Error but returns a JSON error message with format:
// { error: { msg: "error message" } }
func Error(w http.ResponseWriter, msg string, status int) {
  w.WriteHeader(status)
  w.Header().Set("Content-Type", "application/json")

  res := apiError{
    Error: errMsg{msg},
  }
  enc := json.NewEncoder(w)
  if err := enc.Encode(res); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}
