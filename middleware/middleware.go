package middleware

import (
	"crud-fiber/utils"

	"github.com/gofiber/fiber/v2"
)

func Auth(ctx *fiber.Ctx) error {

	token := ctx.Get("x-token")
	if token == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"code":    401,
			"message": "error unauthorized",
		})
	}

	//_, err := utils.VerifyToken(token)
	claims, err := utils.DecodeToken(token)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"code":    401,
			"message": "error unauthorized",
		})
	}

	if role := claims["role"]; role != "admin" {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"code":    fiber.StatusForbidden,
			"message": "forbidden access",
		})
	}
	ctx.Locals("userInfo", claims)
	//ctx.Locals("role", claims["role"])
	return ctx.Next()
}
