package routes

import (
	"crud-fiber/config"
	"crud-fiber/controllers"
	"crud-fiber/middleware"

	"github.com/gofiber/fiber/v2"
)

func InitRoutes(r *fiber.App) {
	r.Static("/images", config.RootPath+"/public/assets/images")
	r.Static("/files", config.RootPath+"/public/assets/files")

	
	r.Post("/login", controllers.Login)
	r.Get("/users", controllers.GetUserAll)
	r.Post("/user", controllers.CreateUser)
	r.Get("/user/:id", middleware.Auth, controllers.UserGetById)
	r.Put("/user/:id", controllers.UserUpdate)
	r.Put("/user/:id/update-email", controllers.UserUpdateEmail)
	r.Delete("/user/:id", controllers.UserDelete)
}
