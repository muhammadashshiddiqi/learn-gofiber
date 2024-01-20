package routes

import (
	"crud-fiber/controllers"

	"github.com/gofiber/fiber/v2"
)

func InitRoutes(r *fiber.App) {
	r.Get("/users", controllers.GetUserAll)
	r.Post("/user", controllers.CreateUser)
	r.Get("/user/:id", controllers.UserGetById)
	r.Put("/user/:id", controllers.UserUpdate)
	r.Put("/user/:id/update-email", controllers.UserUpdateEmail)
	r.Delete("/user/:id", controllers.UserDelete)
}
