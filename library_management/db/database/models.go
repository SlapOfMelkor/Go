// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package database

import (
	"time"
)

type Book struct {
	ID            int32  `json:"id"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	PublishedDate string `json:"published_date"`
}

type Borrow struct {
	ID         int32     `json:"id"`
	UserID     int32     `json:"user_id"`
	BookID     int32     `json:"book_id"`
	Status     string    `json:"status"`
	BorrowDate time.Time `json:"borrow_date"`
	ReturnDate time.Time `json:"return_date"`
}

type User struct {
	ID       int32  `json:"id"`
	Rol      string `json:"rol"`
	Username string `json:"username"`
	Pasword  string `json:"pasword"`
}
