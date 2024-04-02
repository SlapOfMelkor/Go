package controller

import (
	"library_management/db"
	"library_management/db/database"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(context *fiber.Ctx) error {
	usercreate := new(database.CreateUserParams)
	err := context.BodyParser(&usercreate)
	if err != nil {
		return context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"Message": "Request Failed"})
	}
	query := database.New(db.Db)
	contextt := context.Context()
	hashedpass := HashPassword(usercreate.Pasword)
	usercreate.Pasword = hashedpass
	_, err = query.CreateUser(contextt, *usercreate)

	if err != nil {

		return context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"Message": "User could not created"})

	}
	return context.Status(http.StatusOK).JSON(&fiber.Map{"Message": "User succesfully created"})

}
func GetUsers(context *fiber.Ctx) error {

	query := database.New(db.Db)
	contextt := context.Context()
	users, err := query.GetUsers(contextt)
	if err != nil {
		return context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"Message": "Users Could not returned"})

	}
	return context.Status(http.StatusOK).JSON(users)

}
func GetUser(context *fiber.Ctx) error {
	id, _ := strconv.Atoi(context.Params("id"))
	query := database.New(db.Db)
	contextt := context.Context()
	users, err := query.GetUser(contextt, int32(id))
	if err != nil {
		return context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"Message": "Users Could not returned"})

	}
	return context.Status(http.StatusOK).JSON(users)
}
func DeleteUser(context *fiber.Ctx) error {
	id, _ := strconv.Atoi(context.Params("id"))
	query := database.New(db.Db)
	contextt := context.Context()
	err := query.DeleteUser(contextt, int32(id))
	if err != nil {
		return context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"Message": "Users Could not returned"})

	}
	return context.Status(http.StatusOK).JSON("user succesfully deleted")
}

func UpdateUser(context *fiber.Ctx) error {
	id, _ := strconv.Atoi(context.Params("id"))
	query := database.New(db.Db)
	userupdate := new(database.UpdateUserParams)
	err := context.BodyParser(&userupdate)
	if err != nil {
		return context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"Message": "Request Failed"})
	}
	contextt := context.Context()
	userupdate.ID = int32(id)
	err = query.UpdateUser(contextt, *userupdate)

	if err != nil {

		return context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"Message": "User could not updated"})

	}
	return context.Status(http.StatusOK).JSON(&fiber.Map{"Message": "User succesfully updated"})
}
func HashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(hashedPassword)
}
