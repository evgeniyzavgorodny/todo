package repository

import (
	"context"
	"github.com/evgeniyzavgorodny/todo/models"
	"github.com/evgeniyzavgorodny/todo/models/dto"
)

type Todo interface {
	Create(ctx context.Context, item dto.CreateDto) (*models.Todo, error)
	Update(ctx context.Context, id uint64, item dto.UpdateDto) (*models.Todo, error)
	Find(ctx context.Context) ([]models.Todo, error)
	Get(ctx context.Context, id uint64) (*models.Todo, error)
	Delete(ctx context.Context, id uint64) error
}
