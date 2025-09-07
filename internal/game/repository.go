package game

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type GameRepository struct {
	dbpool *pgxpool.Pool
	logger *zerolog.Logger
}

func NewGameRepository(dbpool *pgxpool.Pool, customLogger *zerolog.Logger) *GameRepository {
	return &GameRepository{
		dbpool: dbpool,
		logger: customLogger,
	}
}
