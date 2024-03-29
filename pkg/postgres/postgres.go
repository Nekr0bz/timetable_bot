package postgres

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/log/zapadapter"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
	"time"
)

const connAttempts = 5

type Postgres struct {
	Builder squirrel.StatementBuilderType
	Conn    *pgxpool.Pool
}

func New(dbName, dbHost, dbUser string, log *zap.Logger) (_ *Postgres, err error) {
	connString := fmt.Sprintf("postgres://%s@%s/%s?sslmode=disable", dbUser, dbHost, dbName)
	log.Info("Connecting to postgres", zap.String("connString", connString))

	pgConfig, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}
	pgConfig.ConnConfig.Logger = zapadapter.NewLogger(log)

	for i := 0; i < connAttempts; i++ {
		conn, err := pgxpool.ConnectConfig(context.Background(), pgConfig)
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
