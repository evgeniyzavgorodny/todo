package http

import (
	"context"
	"github.com/evgeniyzavgorodny/todo/models/dto"
	"github.com/evgeniyzavgorodny/todo/repository"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

type TodoHTTP struct {
	repo repository.Todo
}

func (h *TodoHTTP) Create(c *fiber.Ctx) error {
	ctx := context.Background()
	var item dto.CreateDto
	err := c.BodyParser(&item)
	if err != nil {
		return err
	}

	res, err := h.repo.Create(ctx, item)
	c.Status(http.StatusCreated)
	c.Set("Content-Type", "application/json")
	return c.JSON(res)
}

func (h *TodoHTTP) Update(c *fiber.Ctx) error {
	ctx := context.Background()
	var item dto.UpdateDto
	idStr := c.Params("id")
	err := c.BodyParser(&item)
	if err != nil {
		return err
	}

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return err
	}

	res, err := h.repo.Update(ctx, id, item)
	if err != nil {
		return err
	}

	c.Set("Content-Type", "application/json")
	return c.JSON(res)
}

func (h *TodoHTTP) Find(c *fiber.Ctx) error {
	ctx := context.Background()

	res, err := h.repo.Find(ctx)
	if err != nil {
		return err
	}

	c.Set("Content-Type", "application/json")
	return c.JSON(res)
}

func (h *TodoHTTP) Delete(c *fiber.Ctx) error {
	ctx := context.Background()
	idStr := c.Params("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return err
	}

	err = h.repo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return c.SendStatus(http.StatusNoContent)
}

func NewTodoDelivery(fiber *fiber.App, repo repository.Todo) {
	handler := &TodoHTTP{
		repo: repo,
	}

	fiber.Post("/tasks", handler.Create)
	fiber.Put("/tasks/:id", handler.Update)
	fiber.Get("/tasks/", handler.Find)
	fiber.Delete("/tasks/:id", handler.Delete)
}
