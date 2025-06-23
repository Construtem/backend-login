package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("No se pudo conectar a la base de datos: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Ping fallido a la base de datos: %v", err)
	}

	log.Println("✅ Conexión a PostgreSQL establecida.")
}
