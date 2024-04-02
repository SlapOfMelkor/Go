package router

import (
	"library_management/controller"
	"library_management/middleware"

	"github.com/gofiber/fiber/v2"
)

func Bookrouter(app *fiber.App) {
	app.Get("/book/getbookbyid/:id", controller.GetBook, middleware.Ifuser)
	rtstring := app.Group("/book", middleware.Ifadmin)
	rtstring.Post("/addbook", controller.AddBook)
	rtstring.Delete("/deletebookbyid/:id", controller.Deletebook)
	rtstring.Put("/updatebookbyid/:id", controller.UpdateBook)
}
