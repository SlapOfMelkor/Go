package main

import (
	"library_management/db"
	"library_management/router"
	"log"

	// Import your generated database package

	"github.com/gofiber/fiber/v2"
	_ "github.com/mattn/go-sqlite3" // Import SQLite driver
)

func main() {
	db.Connectdb()

	app := fiber.New()
	router.Mainrouter(app)
	log.Fatal(app.Listen(":3000"))

}
