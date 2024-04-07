package postgres

import (
	"context"
	logs "effect/internal/logger"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetParamsForDB() string {
	host := "repository"
	port := "5432"
	user := "repository"
	pass := "123"
	dbname := "hezzlPost"
	params := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, dbname)

	return params
}

func Connect(ctx context.Context, params string) *pgxpool.Pool {

	dbpool, err := pgxpool.New(ctx, params)
	if err != nil {
		logs.LogFatal(logs.Logger, "postgres_connector", "Connect", err, err.Error())
	}

	err = dbpool.Ping(ctx)
	if err != nil {
		logs.LogFatal(logs.Logger, "postgres_connector", "Connect", err, err.Error())
	}
	logs.Logger.Info("Connected to db")
	logs.Logger.Debug("db client :", dbpool)

	return dbpool
}
