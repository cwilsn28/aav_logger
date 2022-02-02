package utils

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

/* --- Establish database connection --- */
func OpenDB() *sql.DB {
	newdb,
		err := sql.Open(
		"postgres",
		fmt.Sprintf(
			`host=%s
             port=5432
             dbname=%s
             user=%s
             password=%s
             sslmode=disable`,
			os.Getenv("PGHOST"),
			os.Getenv("PGNAME"),
			os.Getenv("PGUSER"),
			os.Getenv("PGNAME"),
		),
	)
	if err != nil {
		log.Fatalf("Error opening database connection: %s", err.Error())
	}

	newdb.SetMaxIdleConns(100)
	return newdb
}
