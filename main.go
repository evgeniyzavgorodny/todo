package main

import (
	"context"
	"fmt"
	"github.com/evgeniyzavgorodny/todo/delivery/http"
	"github.com/evgeniyzavgorodny/todo/errorHandler"
	"github.com/evgeniyzavgorodny/todo/migrations"
	pgxHandler "github.com/evgeniyzavgorodny/todo/repository/pgx"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	"log"
	"os"
)

func main() {
	ctx := context.Background()
	dbURL := os.Getenv("DATABASE_URL")
	httpAddress := os.Getenv("HTTP_ADDRESS")

	conn, err := pgx.Connect(ctx, dbURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(ctx)

	migrations.Migrations(ctx, conn)

	todoRepo := pgxHandler.NewTodoRepository(conn)
	app := fiber.New(fiber.Config{
		ErrorHandler: errorHandler.ErrorHandler,
	})
	http.NewTodoDelivery(app, todoRepo)

	log.Fatal(app.Listen(httpAddress))
}
