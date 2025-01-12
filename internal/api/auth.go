package api

import (
	"ridhoandhika/backend-api/domain"
	"ridhoandhika/backend-api/dto"
	"ridhoandhika/backend-api/internal/util"
	"strings"

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
		return ctx.SendStatus(util.GetHttpStatus(domain.ErrAuthFailed))
	}

	return ctx.Status(200).JSON(dto.BaseResp{
		ErrorSchema: dto.ErrorSchema{
			ErrorCode: "200",
			ErrorMessage: dto.ErrorMessage{
				Id: "Sukses", // Perbaiki pengejaan dari "Sukes" menjadi "Sukses"
				En: "Success",
			},
		},
		OutputSchema: token,
	})
}

func (a authApi) ValidateToken(ctx *fiber.Ctx) error {
	authHeader := ctx.Get("Authorization")

	// Memparsing token (memisahkan Bearer dan token)
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return ctx.SendStatus(util.GetHttpStatus(domain.ErrAuthFailed))
	}

	token := parts[1]
	user, err := a.userService.ValidateToken(ctx.Context(), token)
	if err != nil {
		// Jika token tidak valid atau error lain
		return ctx.SendStatus(util.GetHttpStatus(domain.ErrAuthFailed))
	}

	return ctx.Status(200).JSON(dto.BaseResp{
		ErrorSchema: dto.ErrorSchema{
			ErrorCode: "200",
			ErrorMessage: dto.ErrorMessage{
				Id: "Sukses", // Perbaiki pengejaan dari "Sukes" menjadi "Sukses"
				En: "Success",
			},
		},
		OutputSchema: user,
	})
}

func (a authApi) Register(ctx *fiber.Ctx) error {
	var req dto.UserRegisterReq
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(200).JSON(dto.BaseResp{
			ErrorSchema: dto.ErrorSchema{
				ErrorCode: "400",
				ErrorMessage: dto.ErrorMessage{
					Id: "Permintaan Tidak Valid",
					En: "Bad Request",
				},
			},
		})
	}

	user, err := a.userService.Register(ctx.Context(), req)
	if err != nil {
		return ctx.SendStatus(util.GetHttpStatus(err))
	}

	return ctx.Status(200).JSON(user)
}
