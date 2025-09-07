package session

import (
	"encoding/json"
	"fmt"
	"sudoku/pkg/sudoku"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/redis/v2"
)

type SessionStorage struct {
	Store *session.Store
}

func NewSession(redisStorage *redis.Storage) *SessionStorage {
	return &SessionStorage{
		Store: session.New(session.Config{
			Storage:    redisStorage,
			Expiration: 24 * time.Hour,
			KeyLookup:  "cookie:session_id",
		}),
	}

}

func (s *SessionStorage) SetActiveSudokuData(c *fiber.Ctx, sudokuGrid sudoku.SudokuGrid) error {
	sess, err := s.GetSession(c)
	if err != nil {
		return fmt.Errorf("failed to get session: %w", err)
	}
	jsonData, err := json.Marshal(sudokuGrid)
	if err != nil {
		return fmt.Errorf("failed to marshal sudoku data: %w", err)
	}
	sess.Set("sudoku_data", string(jsonData))
	if err := sess.Save(); err != nil {
		return fmt.Errorf("failed to save session: %w", err)
	}
	return nil
}

func (s *SessionStorage) GetSession(c *fiber.Ctx) (*session.Session, error) {
	return s.Store.Get(c)
}

func (s *SessionStorage) SaveSession(sess *session.Session) error {
	return sess.Save()
}

func (s *SessionStorage) GetActiveAllSudokuData(c *fiber.Ctx) (*sudoku.SudokuPair, error) {
	sess, err := s.GetSession(c)
	if err != nil || sess == nil {
		return nil, nil 
	}

	rawPuzzle := sess.Get("sudoku_puzzle_data")
	rawInverted := sess.Get("sudoku_inverted_data")
	rawSolution := sess.Get("sudoku_solution_data")

	if rawPuzzle == nil || rawInverted == nil || rawSolution == nil {
		return nil, nil
	}

	strPuzzle, ok1 := rawPuzzle.(string)
	strInverted, ok2 := rawInverted.(string)
	strSolution, ok3 := rawSolution.(string)
	if !ok1 || !ok2 || !ok3 {
		return nil, fmt.Errorf("invalid session data")
	}

	var puzzle, inverted, solution sudoku.SudokuGrid
	if err := json.Unmarshal([]byte(strPuzzle), &puzzle); err != nil {
		return nil, err
	}
	if err := json.Unmarshal([]byte(strInverted), &inverted); err != nil {
		return nil, err
	}
	if err := json.Unmarshal([]byte(strSolution), &solution); err != nil {
		return nil, err
	}

	return &sudoku.SudokuPair{
		Puzzle:   &puzzle,
		Inverted: &inverted,
		Solution: &solution,
	}, nil
}

