package db

import (
  "fmt"
  "log"
  "os"

  "github.com/jinzhu/gorm"
  // Register postgres dialec
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB holds connection to database
var DB *gorm.DB

func init() {
  if DB == nil {
    var err error
    DB, err = gorm.Open("postgres", ConnectionString())
    log.Println("Openning database connection")
    if err != nil {
      log.Fatalf("Couldn't open connection to database. Error: %s", err.Error())
    }
  }
}

// ConnectionString returns a Postgres connection string with information
// obtained from env var
func ConnectionString() string {
  conn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s",
    os.Getenv("DB_HOST"),
    os.Getenv("DB_PORT"),
    os.Getenv("DB_USERNAME"),
    os.Getenv("DB_NAME"),
    os.Getenv("DB_SSL_MODE"),
  )

  if pwd := os.Getenv("DB_PASSWORD"); pwd != "" {
    conn = fmt.Sprintf("%s password=%s", conn, pwd)
  }

  return conn
}
