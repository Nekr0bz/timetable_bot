package postgres

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

const connAttempts = 10

type Postgres struct {
	Builder squirrel.StatementBuilderType
	Conn    *pgxpool.Pool
}

func New(dbName, dbUser string) (_ *Postgres, err error) {
	var conn *pgxpool.Pool

	for i := 0; i < connAttempts; i++ {
		conn, err = pgxpool.Connect(context.Background(), "postgres://"+dbUser+"@db/"+dbName+"?sslmode=disable")
		if err != nil {
			time.Sleep(1 * time.Second)
			continue
		}

		return &Postgres{
			Conn:    conn,
			Builder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
		}, nil
	}

	return nil, err
}
