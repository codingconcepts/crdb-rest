package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	db, err := pgxpool.New(context.Background(), "postgres://root@localhost:26257/defaultdb?sslmode=disable")
	if err != nil {
		log.Fatalf("error connecting to db: %v", err)
	}
	defer db.Close()

	router := fiber.New()
	router.Get("/v1/todos", getTodos(db))
	router.Get("/v1/todos/:id", getTodo(db))
	router.Post("/v1/todos", createTodo(db))
	router.Delete("/v1/todos/:id", deleteTodo(db))

	log.Fatal(router.Listen(":8080"))
}

func getTodos(db *pgxpool.Pool) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		return nil
	}
}

func getTodo(db *pgxpool.Pool) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		return nil
	}
}

func createTodo(db *pgxpool.Pool) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		return nil
	}
}

func deleteTodo(db *pgxpool.Pool) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		return nil
	}
}
