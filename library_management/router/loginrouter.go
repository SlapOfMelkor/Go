package router

import (
	"library_management/controller"

	"github.com/gofiber/fiber/v2"
)

func LoginRouter(app *fiber.App) {

	app.Post("/api/login", controller.Login)
}
