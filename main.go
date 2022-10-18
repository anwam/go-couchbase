package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type CreateUserDto struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}

type ToInsertDto struct {
	CreateUserDto
	Id string `json:"id"`
}

func main() {
	log.Println("### getting started ###")

	app := fiber.New()
	db := NewStore("dev", "aa111111")
	bucket := db.cb.Bucket("subscribe_dev")
	col := bucket.DefaultScope().Collection("subscribe_profile")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/user", func(c *fiber.Ctx) error {
		dto := new(CreateUserDto)
		c.BodyParser(dto)
		id := uuid.NewString()
		insertResult, err := col.Insert(id, &ToInsertDto{
			CreateUserDto: *dto,
			Id:            id,
		}, nil)
		if err != nil {
			log.Println("insert failed " + err.Error())
			return c.Status(400).JSON(map[string]interface{}{
				"message": "insert failed.",
				"error":   err,
			})
		}
		return c.Status(201).JSON(map[string]string{
			"message": "created",
			"result":  fmt.Sprintf("%+v", insertResult),
		})
	})

	app.Listen(":3000")
}
