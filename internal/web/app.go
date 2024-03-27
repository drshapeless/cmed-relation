package web

import (
	"cmed-relation/internal/logger"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Application struct {
	Port   int
	DB     *pgxpool.Pool
	WG     sync.WaitGroup
	Logger *logger.Logger
}
