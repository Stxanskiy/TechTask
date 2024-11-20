package repo

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

func ConnectDB(dsn string) (*pgxpool.Pool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(ctx); err != nil {
		return nil, err

	}
	return db, err
}
