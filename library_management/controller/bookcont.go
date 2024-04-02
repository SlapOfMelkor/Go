package controller

import (
	"library_management/db"
	"library_management/db/database"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func AddBook(context *fiber.Ctx) error {
	bookcreate := new(database.AddBookParams)
	err := context.BodyParser(&bookcreate)
	if err != nil {
		return context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"Message": "Request Failed"})
	}
	query := database.New(db.Db)
	contextt := context.Context()
	_, err = query.AddBook(contextt, *bookcreate)

	if err != nil {

		return err

	}
	return context.Status(http.StatusOK).JSON(&fiber.Map{"Message": "Book succesfully Added"})

}
func GetBook(context *fiber.Ctx) error {
	id, _ := strconv.Atoi(context.Params("id"))
	query := database.New(db.Db)
	contextt := context.Context()
	users, err := query.GetBookByID(contextt, int32(id))
	if err != nil {
		return context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"Message": "Users Could not returned"})

	}
	return context.Status(http.StatusOK).JSON(users)

}
func Deletebook(context *fiber.Ctx) error {
	id, _ := strconv.Atoi(context.Params("id"))
	query := database.New(db.Db)
	contextt := context.Context()
	err := query.DeleteBook(contextt, int32(id))
	if err != nil {
		return context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"Message": "Users Could not returned"})

	}
	return context.Status(http.StatusOK).JSON("user succesfully deleted")
}
func UpdateBook(context *fiber.Ctx) error {
	id, _ := strconv.Atoi(context.Params("id"))
	query := database.New(db.Db)
	bookupdate := new(database.UpdateBookParams)
	err := context.BodyParser(&bookupdate)
	if err != nil {
		return context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"Message": "Request Failed"})
	}
	contextt := context.Context()
	bookupdate.ID = int32(id)
	err = query.UpdateBook(contextt, *bookupdate)

	if err != nil {

		return context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"Message": "User could not updated"})

	}
	return context.Status(http.StatusOK).JSON(&fiber.Map{"Message": "User succesfully updated"})
}
