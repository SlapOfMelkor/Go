package router

import "github.com/gofiber/fiber/v2"

func Mainrouter(app *fiber.App) {
	LoginRouter(app)
	Borrowrouter(app)
	Userrouter(app)
	Bookrouter(app)

}
