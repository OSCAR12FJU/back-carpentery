package postgres

import (
	"database/sql"
	"fmt"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func CreateConnection(dbURI string) (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error al leer el arch .env: %w", err)
	}

	db, err := sql.Open("postgres", dbURI)
	if err != nil {
		return nil, fmt.Errorf("error al conectar a Postgres: %w", err)
	}
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error al conectar desde ping: %w", err)
	}
	fmt.Println("Conectado a Postgres!")
	return db, nil
}
