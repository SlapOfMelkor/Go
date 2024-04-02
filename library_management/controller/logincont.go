package controller

import (
	"library_management/db"
	"library_management/db/database"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"golang.org/x/crypto/bcrypt"
)

var Sessiondata = session.New()

func Login(context *fiber.Ctx) error {

	logindata := new(database.LoginRow)
	err := context.BodyParser(&logindata)
	if err != nil {
		return context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"Message": "Bad request"})

	}
	loginquery := database.New(db.Db)
	logincontext := context.Context()
	var loginterface database.LoginRow
	loginterface, err = loginquery.Login(logincontext, logindata.Username)
	if err != nil {
		return context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"Message": "Username or password not correct"})

	}
	err = bcrypt.CompareHashAndPassword([]byte(loginterface.Pasword), []byte(logindata.Pasword))
	if err != nil {
		return context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"Message": "Username or password not correct"})
	}
	sessiondata, _ := Sessiondata.Get(context)
	sessiondata.Set("rol", loginterface.Rol)
	err = sessiondata.Save()
	if err != nil {
		panic("fatal error")

	}
	return context.Status(http.StatusOK).JSON(&fiber.Map{"Message": "Succesfuly logged in"})
}
