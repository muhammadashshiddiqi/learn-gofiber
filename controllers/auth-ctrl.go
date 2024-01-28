package controllers

import (
	"crud-fiber/database"
	"crud-fiber/models/entity"
	"crud-fiber/models/request"
	"crud-fiber/utils"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Login(ctx *fiber.Ctx) error {
	loginReq := new(request.LoginRequest)
	if err := ctx.BodyParser(loginReq); err != nil {
		return err
	}
	log.Println(loginReq)

	//check request
	validate := validator.New()
	errValidate := validate.Struct(loginReq)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{

			"code":    400,
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	//check avail email
	var user entity.User
	email := database.DB.First(&user, "email=?", loginReq.Email)
	if email.Error != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"code":    401,
			"message": "wrong credential",
		})
	}

	isValid := utils.CheckPassHash(loginReq.Password, user.Password)
	if !isValid {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"code":    401,
			"message": "wrong credential",
		})
	}

	//GENERATE JWT
	claims := jwt.MapClaims{}
	claims["name"] = user.Name
	claims["email"] = user.Email
	claims["address"] = user.Address
	claims["exp"] = time.Now().Add(time.Minute * 24).Unix()

	if user.Email == "sidiq@test.dev" {
		claims["role"] = "admin"
	}else{
		claims["role"] = "user"
	}

	token, errToken := utils.GenerateToken(&claims)
	if errToken != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"code":    401,
			"message": "wrong credential",
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"code":    200,
		"message": "waiting progress",
		"token": token, 
	})
}
