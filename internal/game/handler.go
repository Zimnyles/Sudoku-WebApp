package game

import (
	"encoding/json"
	"fmt"
	"sudoku/pkg/session"
	"sudoku/pkg/sudoku"
	"sudoku/pkg/tadapter"
	"sudoku/web/pages"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type GameHandler struct {
	router  fiber.Router
	logger  *zerolog.Logger
	store   *session.SessionStorage
	service IGameService
}

type IGameService interface {
	NewGame(difficult int, c *fiber.Ctx) error
}

func NewGameHandler(router fiber.Router, logger *zerolog.Logger, service IGameService, store *session.SessionStorage) {
	h := &GameHandler{
		router:  router,
		logger:  logger,
		store:   store,
		service: service,
	}

	h.router.Get("/", h.game)

	h.router.Post("/api/cell", h.apiCellRequest)

	//for dev tests
	h.router.Get("/actual", h.actual)
	h.router.Get("/test", h.test)

}

func (h *GameHandler) apiCellRequest(c *fiber.Ctx) error {
	req := new(CellRequest)
	if err := c.BodyParser(req); err != nil {
		h.logger.Error().Err(err).Msg("Failed to parse cell request")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	grids, err := h.store.GetActiveAllSudokuData(c)
	if err != nil {
		h.logger.Error().Err(err).Msg("Failed to get active sudoku data")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "session error",
		})
	}

	if req.IsCorrect {
		if grids.Solution.Grid[req.Row][req.Col].Value == req.Value {
			grids.Puzzle.Grid[req.Row][req.Col].Value = req.Value
			if err := h.store.SetActiveSudokuData(c, *grids.Puzzle); err != nil {
				h.logger.Error().Err(err).Msg("Failed to save sudoku data to session")
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "failed to save",
				})
			}
		}

		return c.JSON(fiber.Map{
			"row":   req.Row,
			"col":   req.Col,
			"value": req.Value,
			"ok":    true,
		})
	}

	grids.Empty.Grid[req.Row][req.Col].Value = req.Value
	if err := h.store.SetEmptySudokuData(c, *grids.Empty); err != nil {
		h.logger.Error().Err(err).Msg("Failed to save sudoku data to session")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to save",
		})
	}
	if err := h.store.SetFailData(c); err != nil {
		h.logger.Error().Err(err).Msg("Failed to save fail data to session")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to save",
		})
	}

	return c.JSON(fiber.Map{
		"row":   req.Row,
		"col":   req.Col,
		"value": req.Value,
		"ok":    true,
	})

}

func (h *GameHandler) game(c *fiber.Ctx) error {
	data, err := h.store.GetActiveAllSudokuData(c)
	if err != nil {
		h.logger.Error().Err(err).Msg("Failed to get session data")
		return c.Status(fiber.StatusInternalServerError).SendString("Server error")
	}

	if data == nil {
		if err := h.service.NewGame(hard, c); err != nil {
			h.logger.Error().Err(err).Msg("Failed to create new game")
			return c.Status(fiber.StatusInternalServerError).SendString("Server error")
		}

		data, err = h.store.GetActiveAllSudokuData(c)
		if err != nil || data == nil {
			h.logger.Error().Err(err).Msg("Failed to get newly created game")
			return c.Status(fiber.StatusInternalServerError).SendString("Server error")
		}

		puzzleJSON, _ := json.Marshal(data.Puzzle.Grid)
		solutionJSON, _ := json.Marshal(data.Solution.Grid)

		templateData := templateData{
			Grids:        data,
			PuzzleJSON:   string(puzzleJSON),
			SolutionJSON: string(solutionJSON),
			Fails:        data.Fails,
		}

		data.PrintPair()
		fmt.Println("fails: ", data.Fails)
		comp := pages.GamePage(templateData)
		return tadapter.Render(c, comp, 200)
	}

	puzzleJSON, _ := json.Marshal(data.Puzzle.Grid)
	solutionJSON, _ := json.Marshal(data.Solution.Grid)

	templateData := templateData{
		Grids:        data,
		Fails:        data.Fails,
		PuzzleJSON:   string(puzzleJSON),
		SolutionJSON: string(solutionJSON),
	}

	data.PrintPair()
	fmt.Println("fails: ", data.Fails)
	comp := pages.GamePage(templateData)
	return tadapter.Render(c, comp, 200)
}

func (h *GameHandler) actual(c *fiber.Ctx) error {
	data, err := h.store.GetActiveAllSudokuData(c)
	if err != nil {
		h.logger.Error().Err(err).Msg("error")
		return c.Status(fiber.StatusInternalServerError).SendString("error")
	}
	fmt.Println(data)
	return nil
}

func (h *GameHandler) test(c *fiber.Ctx) error {
	data := sudoku.NewSudokuPair(40)

	data.PrintPair()

	return nil
}
