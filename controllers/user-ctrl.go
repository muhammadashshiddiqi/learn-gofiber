package controllers

import (
	"crud-fiber/database"
	"crud-fiber/models/entity"
	"crud-fiber/models/repository"
	"crud-fiber/models/request"
	"crud-fiber/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func GetUserAll(ctx *fiber.Ctx) error {
	var users []entity.User
	//result := database.DB.Debug().Find(&users)
	result := database.DB.Find(&users)

	if result.Error != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    400,
			"message": result.Error,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"code":   200,
			"status": "success",
			"data":   users,
		})
}

func CreateUser(ctx *fiber.Ctx) error {
	user := new(request.UserCreateRequest)
	if err := ctx.BodyParser(user); err != nil {
		return err
	}

	validate := validator.New()
	errValidate := validate.Struct(user)
	if errValidate != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    400,
			"message": errValidate.Error(),
		})
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    500,
			"message": err,
		})
	}

	newUser := entity.User{
		Name:     user.Name,
		Email:    user.Email,
		Address:  user.Address,
		Phone:    user.Phone,
		Password: hashedPassword,
	}

	result := database.DB.Create(&newUser)
	if result.Error != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    400,
			"message": result.Error,
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"code":    200,
		"message": "successfully, user created",
		"data":    newUser,
	})
}

func UserGetById(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")
	var user entity.User
	result := database.DB.First(&user, "id=?", userId)
	if result.Error != nil {
		return ctx.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"code":    400,
			"message": "user not found",
		})
	}

	userRes := repository.User{
		Name:      user.Name,
		Phone:     user.Phone,
		Email:     user.Email,
		Address:   user.Address,
		CreatedAt: user.CreatedAt,
		DeletedAt: nil,
	}

	return ctx.Status(fiber.StatusCreated).JSON(
		fiber.Map{
			"code":   200,
			"status": "success",
			"data":   userRes,
		})
}

func UserUpdate(ctx *fiber.Ctx) error {
	userRequest := new(request.UserUpdateRequest)
	if err := ctx.BodyParser(userRequest); err != nil {
		return ctx.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"code":    400,
			"message": "bad request",
		})
	}

	userId := ctx.Params("id")
	var user entity.User
	result := database.DB.First(&user, "id=?", userId)
	if result.Error != nil {
		return ctx.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"code":    400,
			"message": "user not found",
		})
	}

	if userRequest.Name != "" {
		user.Name = userRequest.Name
	}
	user.Address = userRequest.Address
	user.Phone = userRequest.Phone

	resultUpdate := database.DB.Save(&user)
	if resultUpdate.Error != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"code":    500,
			"message": "internal server error",
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"code":    200,
		"message": "successfully, user updated",
		"data":    user,
	})
}

func UserUpdateEmail(ctx *fiber.Ctx) error {
	userRequest := new(request.UserEmailRequest)
	if err := ctx.BodyParser(userRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"code":    400,
			"message": "bad request",
		})
	}

	userId := ctx.Params("id")
	var user entity.User
	var isEmailUserExist entity.User
	result := database.DB.First(&user, "id=?", userId)
	if result.Error != nil {
		return ctx.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"code":    400,
			"message": "user not found",
		})
	}

	errCheckEmail := database.DB.First(&isEmailUserExist, "email = ?", userRequest.Email).Error
	if errCheckEmail == nil {
		return ctx.Status(402).JSON(fiber.Map{
			"code":    402,
			"message": "email already used.",
		})
	}

	user.Email = userRequest.Email
	resultUpdate := database.DB.Save(&user)
	if resultUpdate.Error != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"code":    500,
			"message": "internal server error",
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"code":    200,
		"message": "successfully, email updated",
		"data":    user,
	})
}
func UserDelete(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")
	var user entity.User
	result := database.DB.Debug().First(&user, "id=?", userId)
	if result.Error != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"code":    400,
			"message": "user not found",
		})
	}

	resultDelete := database.DB.Debug().Delete(&user)
	if resultDelete.Error != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"code":    500,
			"message": "internal server error",
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"code":    200,
		"message": "successfully, user name " + user.Name + "has been deleted",
	})
}
