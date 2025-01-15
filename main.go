package main

import (
	"github.com/Minyeng/go-auth-jwt/auth"
	"github.com/Minyeng/go-auth-jwt/config"
	"github.com/Minyeng/go-auth-jwt/controller"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	db, err := config.CreateDbEngine()
	if err != nil {
		panic(err)
	}

	controller.AuthController(app, db)

	private := app.Group("/private")
	private.Use(auth.JwtWare())
	private.Get("/", func(c *fiber.Ctx) error {

		return c.JSON(fiber.Map{"success": true, "path": "private"})
	})

	app.Listen(":9000")
}
