package database

import (
	"context"
	"fmt"

	"github.com/dinizgab/buildco-api/config"
	"github.com/jackc/pgx/v5"
)

func New(config config.DBConfig) (*pgx.Conn, error) {
    ctx := context.Background()
    dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", config.UserName, config.Password, config.Host, config.Port, config.DBName)

	db, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		return nil, err
	}
    
    defer db.Close(ctx)

	err = db.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
