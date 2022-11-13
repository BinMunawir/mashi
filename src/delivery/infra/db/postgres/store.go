package postgres

import (
	"context"
	"fmt"
	"log"
	"os"

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

	err = pool.Ping(context.Background())
	if err != nil {
		log.Println(err)
		return PostgresStore{}, fmt.Errorf("unable to create connection pool: %v", err)
	}

	return PostgresStore{pool}, nil
}

func (s PostgresStore) SaveInvoice(data map[string]interface{}) {
	id := data["id"].(string)
	title := data["title"].(string)
	query := "INSERT INTO invoices(id, title) values('" + id + "','" + title + "');"
	_, err := s.pool.Query(context.Background(), query)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
}
