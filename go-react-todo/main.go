package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type Todo struct {
	ID          int    `json:"id"`
	Completed   bool   `json:"completed"`
	Description string `json:"description"`
}

func main() {
	fmt.Println("Hello World!")
	app := fiber.New()

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")

	todos := []Todo{}

	app.Get("/api/todos", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(todos)
	})

	// Create a TODO
	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := Todo{}
		if err := c.BodyParser(&todo); err != nil {
			return err
		}

		if todo.Description == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Todo Description is Required"})
		}

		todo.ID = len(todos) + 1
		todos = append(todos, todo)

		return c.Status(201).JSON(todo)
	})

	//Update a TODO
	app.Patch("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos[i].Completed = true
				return c.Status(200).JSON(todos[i])
			}
		}

		return c.Status(404).JSON(fiber.Map{"error": "TODO not found with the given id"})
	})

	//Delete a TODO
	app.Delete("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos = append(todos[:i], todos[i+1:]...)
				return c.Status(200).JSON(fiber.Map{"success": true})
			}
		}
		return c.Status(404).JSON(fiber.Map{"error": "TODO not found with the given id"})
	})

	log.Fatal(app.Listen(":" + PORT))
}
