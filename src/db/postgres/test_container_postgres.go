package db_postgres

import (
	"context"
	"errors"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

var ctx context.Context

func UpTestContainerPostgres() (*postgres.PostgresContainer, error) {
	ctx = context.Background()
	pgContainer, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("postgres:15.3-alpine"),
		postgres.WithDatabase("test-db"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		return nil, errors.New("error to up postgres test container")
	}
	return pgContainer, nil
}

func ReturnTestContainerConnectionString(pgContainer *postgres.PostgresContainer) (*string, error) {
	connStr, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		return nil, errors.New("error to take connection string of pg container")
	}
	return &connStr, nil
}

func DownTestContainerPostgres(pgContainer *postgres.PostgresContainer) error {
	if err := pgContainer.Terminate(ctx); err != nil {
		return errors.New("failed to terminate pgContainer")
	}
	return nil
}
