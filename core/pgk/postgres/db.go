package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewDBPool(DBString string) (*pgxpool.Pool, error) {
	ctx := context.Background()
    connString := DBString
    config, err := pgxpool.ParseConfig(connString)
    if err != nil {
        fmt.Println(err)
        return nil, fmt.Errorf("parsing connection string: %w", err)
    }


    // Optional: Configure pool
    config.MaxConns = 10
    config.MinConns = 2

    pool, err := pgxpool.NewWithConfig(ctx, config)
    if err != nil {
        fmt.Println(err)
        return nil, fmt.Errorf("connecting to database: %w", err)
    }

    return pool, nil
}