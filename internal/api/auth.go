package api

import (
	"ridhoandhika/backend-api/domain"
	"ridhoandhika/backend-api/dto"
	"ridhoandhika/backend-api/internal/util"

	"github.com/gofiber/fiber/v2"
)

type authApi struct {
	userService domain.UserService
}

func Auth(app *fiber.Group, userService domain.UserService, authMid fiber.Handler) {
	handler := authApi{
		userService: userService,
	}
	app.Post("auth/login", handler.GenerateToken)
	app.Get("auth/refresh", authMid, handler.ValidateToken)
	app.Post("auth/register", handler.Register)
}

func (a authApi) GenerateToken(ctx *fiber.Ctx) error {
	var req dto.AuthReq
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(400)
	}

	token, err := a.userService.Authenticate(ctx.Context(), req)
	if err != nil {
		return ctx.SendStatus(util.GetHttpStatus(err))
	}

	return ctx.Status(200).JSON(token)
}

func (a authApi) ValidateToken(ctx *fiber.Ctx) error {
	user := ctx.Locals("x-user")

	return ctx.Status(200).JSON(user)
}

func (a authApi) Register(ctx *fiber.Ctx) error {
	var req dto.UserRegisterReq
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(400)
	}

	user, err := a.userService.Register(ctx.Context(), req)
	if err != nil {
		return ctx.SendStatus(util.GetHttpStatus(err))
	}

	return ctx.Status(200).JSON(user)
}
