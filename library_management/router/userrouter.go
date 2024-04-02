package router

import (
	"library_management/controller"
	

	"github.com/gofiber/fiber/v2"
)

func Userrouter(app *fiber.App) {
	rtstring := app.Group("/user")
	rtstring.Get("/getallusers", controller.GetUsers)
	rtstring.Get("/getuserbyid/:id", controller.GetUser)
	rtstring.Post("/createuser", controller.CreateUser)
	rtstring.Delete("/deleteuser/:id", controller.DeleteUser)
	rtstring.Put("/updateuser/:id", controller.UpdateUser)
}
