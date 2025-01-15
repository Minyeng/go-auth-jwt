package controller

import (
	"github.com/Minyeng/go-auth-jwt/model"
	"github.com/Minyeng/go-auth-jwt/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AuthController(app *fiber.App, db *gorm.DB) {
	app.Post("/signup", func(c *fiber.Ctx) error {
		req := new(model.SignupRequest)
		if err := c.BodyParser(req); err != nil {
			return err
		}

		token, exp, user, err := service.UserSignup(db, req)
		if err != nil {
			return err
		}

		return c.JSON(fiber.Map{"token": token, "exp": exp, "user": user})
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		req := new(model.LoginRequest)
		if err := c.BodyParser(req); err != nil {
			return err
		}

		if req.Email == "" || req.Password == "" {
			return fiber.NewError(fiber.StatusBadRequest, "invalid login credentials")
		}

		token, exp, user, err := service.UserLogin(db, req)
		if err != nil {
			return err
		}

		return c.JSON(fiber.Map{"token": token, "exp": exp, "user": user})
	})
}
