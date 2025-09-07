package game

import (
	"encoding/json"
	"sudoku/pkg/session"
	"sudoku/pkg/sudoku"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/storage/redis/v2"
	"github.com/rs/zerolog"
)

type GameService struct {
	logger     *zerolog.Logger
	repository IGameRepository
	store      *session.SessionStorage
	redis      *redis.Storage
}

type IGameRepository interface{}

func NewGameService(logger *zerolog.Logger, repository IGameRepository, store *session.SessionStorage, redis *redis.Storage) *GameService {
	return &GameService{
		logger:     logger,
		repository: repository,
		store:      store,
		redis:      redis,
	}
}

func (s *GameService) NewGame(difficult int, c *fiber.Ctx) error {
	sess, err := s.store.GetSession(c)
	if err != nil {
		s.logger.Error().Err(err).Msg("Failed to get session")
		return c.Status(fiber.StatusInternalServerError).SendString("Session error")
	}

	grids := s.newGrids(difficult)

	jsonSudokuPuzzle, err := json.Marshal(grids.Puzzle)
	if err != nil {
		s.logger.Error().Err(err).Msg("Failed to marshal jsonSudokuPuzzle data")
		return c.Status(fiber.StatusInternalServerError).SendString("InternalServerError")
	}
	sess.Set("sudoku_puzzle_data", string(jsonSudokuPuzzle))

	jsonSudokuInverted, err := json.Marshal(grids.Inverted)
	if err != nil {
		s.logger.Error().Err(err).Msg("Failed to marshal jsonSudokuInverted data")
		return c.Status(fiber.StatusInternalServerError).SendString("InternalServerError")
	}
	sess.Set("sudoku_inverted_data", string(jsonSudokuInverted))

	jsonSudokuSolution, err := json.Marshal(grids.Solution)
	if err != nil {
		s.logger.Error().Err(err).Msg("Failed to marshal jsonSudokuInverted data")
		return c.Status(fiber.StatusInternalServerError).SendString("InternalServerError")
	}
	sess.Set("sudoku_solution_data", string(jsonSudokuSolution))

	jsonSudokuEmpty, err := json.Marshal(grids.Empty)
	if err != nil {
		s.logger.Error().Err(err).Msg("Failed to marshal jsonSudokuInverted data")
		return c.Status(fiber.StatusInternalServerError).SendString("InternalServerError")
	}
	sess.Set("sudoku_empty_data", string(jsonSudokuEmpty))

	sess.Set("fails", 0)

	s.store.SaveSession(sess)

	return nil
}

func (s *GameService) newGrids(difficulty int) sudoku.SudokuPair {
	var emptyCells int

	switch difficulty {
	case superEasy, easy, middle, hard, superHard, extreme:
		emptyCells = difficulty
	default:
		emptyCells = easy
	}

	pair := sudoku.NewSudokuPair(emptyCells)
	return *pair
}
