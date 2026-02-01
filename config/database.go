package config

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func ConnectDatabase() *sql.DB {
	dsn := "postgresql://postgres:postgres@localhost:5432/mydb?sslmode=disable"

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("✅ connected to database")
	return db
}
