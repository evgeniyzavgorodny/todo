package pgxHandler

import (
	"context"
	"errors"
	"github.com/evgeniyzavgorodny/todo/errorHandler"
	"github.com/evgeniyzavgorodny/todo/models"
	"github.com/evgeniyzavgorodny/todo/models/dto"
	"github.com/evgeniyzavgorodny/todo/repository"
	"github.com/jackc/pgx/v5"
)

type TodoRepo struct {
	pgx *pgx.Conn
}

func (t *TodoRepo) Create(ctx context.Context, item dto.CreateDto) (*models.Todo, error) {
	var todo models.Todo
	query := `
        INSERT INTO todo (title, description) 
        VALUES ($1, $2) 
        RETURNING id, title, description, status, created_at, updated_at
    `

	err := t.pgx.QueryRow(ctx, query, item.Title, item.Description).Scan(
		&todo.Id,
		&todo.Title,
		&todo.Description,
		&todo.Status,
		&todo.CreatedAt,
		&todo.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func (t *TodoRepo) Update(ctx context.Context, id uint64, item dto.UpdateDto) (*models.Todo, error) {
	existTodo, err := t.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	query := `UPDATE todo SET title = $1, description = $2, status = $3 WHERE id = $4 RETURNING id, title, description, status, created_at, updated_at`

	if item.Title != nil {
		existTodo.Title = *item.Title
	}
	if item.Description != nil {
		existTodo.Description = item.Description
	}
	if item.Status != nil {
		switch *item.Status {
		case models.New, models.Done, models.InProgress:
			existTodo.Status = *item.Status
		default:
			return nil, errors.New("Invalid status")
		}

	}

	err = t.pgx.QueryRow(ctx, query, existTodo.Title, existTodo.Description, existTodo.Status, existTodo.Id).Scan(
		&existTodo.Id,
		&existTodo.Title,
		&existTodo.Description,
		&existTodo.Status,
		&existTodo.CreatedAt,
		&existTodo.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return existTodo, err
}

func (t *TodoRepo) Find(ctx context.Context) ([]models.Todo, error) {
	todos := []models.Todo{}

	query := `SELECT * FROM todo`
	rows, err := t.pgx.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var todo models.Todo

		err := rows.Scan(
			&todo.Id,
			&todo.Title,
			&todo.Description,
			&todo.Status,
			&todo.CreatedAt,
			&todo.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		todos = append(todos, todo)
	}

	return todos, err
}

func (t *TodoRepo) Get(ctx context.Context, id uint64) (*models.Todo, error) {
	var todo models.Todo
	query := `SELECT id, title, description, status, created_at, updated_at FROM todo WHERE id = $1`

	err := t.pgx.QueryRow(ctx, query, id).Scan(
		&todo.Id,
		&todo.Title,
		&todo.Description,
		&todo.Status,
		&todo.CreatedAt,
		&todo.UpdatedAt,
	)

	if err == pgx.ErrNoRows {
		return nil, errorHandler.NewNotFoundError("Todo not found")
	} else if err != nil {
		return nil, err
	}

	return &todo, err
}

func (t *TodoRepo) Delete(ctx context.Context, id uint64) error {
	exist, err := t.Get(ctx, id)
	if err != nil {
		return err
	}

	query := `DELETE FROM todo WHERE id = $1`

	commandTag, err := t.pgx.Exec(ctx, query, exist.Id)
	if err != nil {
		return errors.New("Error deletion")
	}

	if commandTag.RowsAffected() == 0 {
		return errorHandler.NewNotFoundError("Todo not found")
	}

	return nil
}

func NewTodoRepository(pgx *pgx.Conn) repository.Todo {
	s := TodoRepo{
		pgx,
	}

	return &s
}
