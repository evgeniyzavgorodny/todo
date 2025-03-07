package dto

import "github.com/evgeniyzavgorodny/todo/models"

type UpdateDto struct {
	Title       *string                `json:"title"`
	Description *string                `json:"description"`
	Status      *models.TodoStatusType `json:"status"`
}
