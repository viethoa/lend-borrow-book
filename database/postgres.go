package database

import (
	"fmt"
	"lend-borrow-book/env"

	"database/sql"
  _ "github.com/lib/pq"
)

func New() {
	return sql.Open("postgres", "user=VietHoa dbname=lendborrowbook sslmode=disable")
}

func getConnectionString() string {
	connectionStringTemplate := "host=%s port=%s sslmode=%s user=%s password='%s' dbname=%s "
	return fmt.Sprintf(connectionStringTemplate,
		env.Get("DB_HOST"),
		env.Get("DB_PORT"),
		env.Get("DB_SSL_MODE"),
		env.Get("DB_USERNAME"),
		env.Get("DB_PASSWORD"),
		env.Get("DB_NAME"))
}