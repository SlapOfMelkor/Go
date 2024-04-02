package middleware

import (
	"library_management/controller"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func Ifuser(context *fiber.Ctx) error {

	check, _ := controller.Sessiondata.Get(context)

	userrole := check.Get("rol")
	if userrole == "user" || userrole == "admin" {
		return context.Next()
	} else {
		return context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"Message": "You are not user"})
	}
}

func Ifadmin(context *fiber.Ctx) error {

	check, _ := controller.Sessiondata.Get(context)

	userrole := check.Get("rol")
	if userrole == "admin" {
		return context.Next()
	} else {
		return context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"Message": "You are not admin"})
	}
}
