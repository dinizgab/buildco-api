package containers

import (
	"context"
	"database/sql"
	"errors"
	"os"
	"path/filepath"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

type PostgresContainer struct {
	*postgres.PostgresContainer
	DBConn *sql.DB
}

func CreatePostgresContainer(ctx context.Context) (*PostgresContainer, error) {
	containerImage := "postgres:latest"

	pgContainer, err := postgres.RunContainer(ctx,
		testcontainers.WithImage(containerImage),
		postgres.WithDatabase("buildco-test"),
		postgres.WithPassword("test-password"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").WithOccurrence(2).WithStartupTimeout(5*time.Second),
		),
	)
	if err != nil {
		return nil, err
	}

	connStr, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		return nil, err
	}

	migrationDirPath, err := searchMigrationDir()
	if err != nil {
		return nil, err
	}

	db, err := goose.OpenDBWithDriver("postgres", connStr)
	if err != nil {
		return nil, err
	}
  
  err = goose.UpContext(ctx, db, migrationDirPath)
	if err != nil {
		return nil, err
	}

	return &PostgresContainer{
		PostgresContainer: pgContainer,
		DBConn:            db,
	}, nil
}

func searchMigrationDir() (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		goModPath := filepath.Join(currentDir, "go.mod")
		if _, err := os.Stat(goModPath); err == nil {
			break
		}

		parent := filepath.Dir(currentDir)
		if parent == currentDir {
			return "", errors.New("go.mod not found")
		}
		currentDir = parent
	}
	migrationDir := filepath.Join(currentDir, "migrations")

	return migrationDir, nil
}
