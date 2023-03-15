package main

import (
	"context"
	"fmt"
	"log"

	"github.com/codingconcepts/crdb-rest/model"
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

	log.Fatal(router.Listen(":3000"))
}

func getTodos(db *pgxpool.Pool) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		page := ctx.QueryInt("page", 1)
		perPage := ctx.QueryInt("per_page", 3)

		const stmt = `SELECT id, title FROM todo OFFSET $1 LIMIT $2`

		offset := (page - 1) * perPage
		rows, err := db.Query(ctx.Context(), stmt, offset, perPage)
		if err != nil {
			return fmt.Errorf("fetching rows: %w", err)
		}

		var todos []model.Todo
		for rows.Next() {
			t := model.Todo{}
			if err = rows.Scan(&t.ID, &t.Title); err != nil {
				return fmt.Errorf("reading todo: %w", err)
			}
			todos = append(todos, t)
		}

		return ctx.JSON(todos)
	}
}

func getTodo(db *pgxpool.Pool) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")

		const stmt = `SELECT title FROM todo WHERE id = $1 LIMIT 1`

		rows := db.QueryRow(ctx.Context(), stmt, id)

		t := model.Todo{
			ID: id,
		}
		if err := rows.Scan(&t.Title); err != nil {
			return fmt.Errorf("reading todo: %w", err)
		}

		return ctx.JSON(t)
	}
}

func createTodo(db *pgxpool.Pool) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		var todo model.Todo
		if err := ctx.BodyParser(&todo); err != nil {
			return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
		}

		const stmt = `INSERT INTO todo (title) VALUES ($1) RETURNING id`

		rows := db.QueryRow(ctx.Context(), stmt, todo.Title)

		if err := rows.Scan(&todo.ID); err != nil {
			return fmt.Errorf("reading todo: %w", err)
		}

		return ctx.JSON(todo)
	}
}

func deleteTodo(db *pgxpool.Pool) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")

		const stmt = `DELETE FROM todo WHERE id = $1 LIMIT 1`

		result, err := db.Exec(ctx.Context(), stmt, id)
		if err != nil {
			return fmt.Errorf("deleting todo: %w", err)
		}

		resp := map[string]int64{
			"affected": result.RowsAffected(),
		}

		return ctx.JSON(resp)
	}
}
