package controller

import (
	"library_management/db"
	"library_management/db/database"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

//	type BookBorrowParams struct {
//		BorrowDate sql.NullTime  `json:"borrow_date"`
//		ReturnDate sql.NullTime  `json:"return_date"`
//		UserID     sql.NullInt32 `json:"user_id"`
//		BookID     sql.NullInt32 `json:"book_id"`
//		Status     string        `json:"status"`
//	}
func BorrowCreate(context *fiber.Ctx) error {

	borrowcreate := new(database.BookBorrowParams)
	err := context.BodyParser(&borrowcreate)

	if err != nil {

		return context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"Message": "Borrows Could not created"})

	}
	borrowcreate.BorrowDate = time.Now()
	borrowcreate.ReturnDate = time.Now().AddDate(0, 0, 14)
	borrowcreate.Status = "borrowed"
	query := database.New(db.Db)
	contextt := context.Context()
	_, err = query.BookBorrow(contextt, *borrowcreate)

	if err != nil {
		return err
	}
	return context.Status(http.StatusOK).JSON(&fiber.Map{"Message": "Borrow succesfully Added"})

}
func BorrowHistory(context *fiber.Ctx) error {

	query := database.New(db.Db)
	contextt := context.Context()
	Borrows, err := query.GetBorrowHistory(contextt)
	if err != nil {
		return context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"Message": "Borrows Could not returned"})

	}
	return context.Status(http.StatusOK).JSON(Borrows)

}
