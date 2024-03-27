package main

import (
	"cmed-relation/internal/logger"
	"cmed-relation/internal/web"
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	var app web.Application
	dbURL := os.Getenv("DATABASE_URL")
	app.Port = 9696

	mylogger := logger.NewLogger()
	mylogger.Level = logger.LogLevelTrace

	app.Logger = &mylogger

	app.Logger.Info(fmt.Sprintf("connecting to %s", dbURL))
	dbpool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		app.Logger.Error(fmt.Sprintf("Unable to create connection pool: %v", err))
		os.Exit(1)
	}
	defer dbpool.Close()

	app.DB = dbpool

	err = app.Serve()
	if err != nil {
		app.Logger.Error(fmt.Sprintf("%v", err))
		os.Exit(1)
	}
}
