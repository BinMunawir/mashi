package postgres

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresStore struct {
	pool *pgxpool.Pool
}

func NewPostgresStore(dsn string) (PostgresStore, error) {
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Println(err)
		return PostgresStore{}, fmt.Errorf("unable to create connection pool: %v", err)
	}
	// defer pool.Close()

	pool.Reset()
	err = pool.Ping(context.Background())
	if err != nil {
		log.Println(err)
		return PostgresStore{}, fmt.Errorf("unable to create connection pool: %v", err)
	}

	return PostgresStore{pool}, nil
}
