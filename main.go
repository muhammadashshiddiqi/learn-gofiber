package main

import (
	"crud-fiber/database"
	"crud-fiber/database/migrations"
	"crud-fiber/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	//init db
	database.ConnectDB()
	//init migration
	migrations.Migrate()

	//init route
	app := fiber.New()
	routes.InitRoutes(app)

	app.Listen(":9090")
}
