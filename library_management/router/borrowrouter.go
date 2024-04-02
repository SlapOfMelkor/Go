package router

import (
	"library_management/controller"
	"library_management/middleware"

	"github.com/gofiber/fiber/v2"
)

func Borrowrouter(app *fiber.App) {
	rtstring := app.Group("/borrow", middleware.Ifuser)
	rtstring.Get("/history", controller.BorrowHistory)
	rtstring.Post("/createborrow", controller.BorrowCreate)
}
