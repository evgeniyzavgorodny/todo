package models

import (
	"time"
)

type TodoStatusType string

const (
	New        TodoStatusType = "new"
	InProgress TodoStatusType = "in_progress"
	Done       TodoStatusType = "done"
)

type Todo struct {
	Id          uint64         `json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	Title       string         `json:"title"`
	Description *string        `json:"description,omitempty"`
	Status      TodoStatusType `json:"status"`
}
