package postgres

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/log/zapadapter"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
	"time"
)

const connAttempts = 10

type Postgres struct {
	Builder squirrel.StatementBuilderType
	Conn    *pgxpool.Pool
}

func New(dbName, dbHost, dbUser string) (_ *Postgres, err error) {
	// TODO: вынести логгеры все!!

	var conn *pgxpool.Pool
	l, _ := zap.NewDevelopment()
	logger := zapadapter.NewLogger(l)
	connString := "postgres://" + dbUser + "@" + dbHost + "/" + dbName + "?sslmode=disable"

	pgConfig, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}
	pgConfig.ConnConfig.Logger = logger

	for i := 0; i < connAttempts; i++ {
		conn, err = pgxpool.ConnectConfig(context.Background(), pgConfig)
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
