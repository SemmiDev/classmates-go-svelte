package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Student struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Identifier string `json:"identifier"`
	Email      string `json:"email"`
	ImageURL   string `json:"image_url"`
}

var classmates = []Student {
	{
		ID:         1,
		Name:       "sammi",
		Identifier: "1010101010101",
		Email:      "sammi@gmail.com",
		ImageURL:   "https://raw.githubusercontent.com/SemmiDev/classmates-go-svelte/main/assets/sammi.png",
	},
	{
		ID:         2,
		Name:       "Rahmatul Izzah Annisa",
		Identifier: "2020202020121",
		Email:      "izzah@gmail.com",
		ImageURL:   "https://raw.githubusercontent.com/izzaah/tugas-imk2/main/img/img/izzah.jpeg",
	},
	{
		ID:         3,
		Name:       "Ayatullah Ramadhan Jacoeb",
		Identifier: "3030303030303",
		Email:      "ayatullah@gmail.com",
		ImageURL:   "https://raw.githubusercontent.com/izzaah/tugas-imk2/main/img/img/keluarga.png",
	},
}

func main() {
	app := fiber.New()
	app.Use(cors.New())

	data := struct {
		Data []Student `json:"data"`
	}{
		classmates,
	}

	app.Get("/api/v1/classmates", func(c *fiber.Ctx) error {
		c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
		return c.Status(fiber.StatusOK).JSON(data)
	})
	app.Get("/api/v1/classmates/:id", func(c *fiber.Ctx) error {
		id,_ := c.ParamsInt("id")
		for _ ,v := range classmates {
			if v.ID == id {
				c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
				return c.Status(fiber.StatusOK).JSON(v)
			}
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "classmate not found"})
	})
	app.Listen(":8080")
}
