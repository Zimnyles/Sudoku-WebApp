package main

import (
	"sudoku/config"
	"sudoku/internal/game"
	"sudoku/pkg/database"
	"sudoku/pkg/logger"
	"sudoku/pkg/redisstorage"
	"sudoku/pkg/session"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	config.Init()

	//logger
	loggerConfig := config.NewLogConfig()
	logger := logger.NewLogger(loggerConfig)

	//redis
	redisConfig := config.NewRedisConfig()
	redisStorage := redisstorage.NewRedisStorage(*redisConfig)

	//database
	databaseConfig := config.NewDBConfig()
	databasePool := database.CreateDataBasePool(databaseConfig, logger)
	defer databasePool.Close()

	app := fiber.New()
	app.Static("/public", "./web/public")
	app.Static("/static", "./web/static")

	store := session.NewSession(redisStorage)

	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: logger,
	}))
	app.Use(recover.New())

	gameRepository := game.NewGameRepository(databasePool, logger)

	gameService := game.NewGameService(logger, gameRepository, store, redisStorage)

	game.NewGameHandler(app, logger, gameService, store)

	
	app.Listen(":3030")
}
