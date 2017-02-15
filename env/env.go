package env

import (
  "os"
  "strconv"
)

func Get(name string) string {
  return os.Getenv(name)
}

func GetPort() string {
  port := os.Getenv("PORT")
  if (port == "") {
    port = "8080"
  }
  return port
}

func GetInt(name string) int {
  if value, err := strconv.Atoi(Get(name)); err == nil {
    return value
  }

  return 0
}