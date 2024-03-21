package database

import (
	"database/sql"
	"fmt"

	"github.com/dinizgab/buildco-api/config"
	_ "github.com/jackc/pgx/v5"
)

func New(config config.DBConfig) (*sql.DB, error) {
    dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", config.UserName, config.Password, config.Host, config.Port, config.DBName)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
    
    defer db.Close()

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
