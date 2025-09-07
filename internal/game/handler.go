package game

import (
	"fmt"
	"sudoku/pkg/session"
	"sudoku/pkg/sudoku"

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

//1 Получаем значение, линию, колонку
//2 Берем из редиса актуальную сетку
//3 Подставляем в нужное место значение
//4 Если правильно, то обновляем актуальную сетку и отдаем ответ
//4 Если НЕ правильно, то НЕ обновляем актуальную сетку и отдаем ответ

func (h *GameHandler) apiCellRequest(c *fiber.Ctx) error {
	return nil
}

//1 Проверяем есть ли активная игровая сессия
//2 Если нет создаем новую
//3 Сохраняем в редис
//4 Отдаем страницу

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
	}

	return c.JSON(data)
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

	data.PrintPair() // выведет все три сетки красиво

	return nil
}
