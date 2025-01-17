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
	// Menambahkan anotasi Swagger untuk login
	app.Post("auth/login", handler.GenerateToken)
	// Menambahkan anotasi Swagger untuk refresh
	app.Get("auth/refresh", authMid, handler.ValidateToken)
	// Menambahkan anotasi Swagger untuk refresh
	app.Post("auth/register", handler.Register)
}

// @Summary Generate Token for Authentication
// @Description Authenticate user and generate JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param body body dto.AuthReq true "Login Credentials" // Mendefinisikan parameter yang dikirimkan dalam request body
// @Success 200 {object} dto.BaseResp{output_schema=dto.AuthResp} "JWT Token"
// @Failure 400 {object} dto.ErrorSchema "Invalid Request"
// @Failure 401 {object} dto.ErrorSchema "Authentication Failed"
// @Router /api/auth/login [post]
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

// @Summary Refresh JWT Token
// @Description Refresh the JWT token for a logged-in user
// @Tags auth
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Authorization header string true "Bearer JWT Token" // Menambahkan parameter Authorization di header
// @Success 200 {object} dto.BaseResp{output_schema=dto.AuthResp} "New JWT Token"
// @Failure 401 {object} dto.ErrorSchema "Authentication Failed"
// @Router /api/auth/refresh [get]
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

// @Summary Register a new user
// @Description Register a new user by providing user credentials
// @Tags auth
// @Accept json
// @Produce json
// @Param body body dto.UserRegisterReq true "User Registration Request"
// @Success 200 {object} dto.BaseResp "Registration Success"
// @Failure 400 {object} dto.ErrorSchema "Invalid Request"
// @Failure 409 {object} dto.ErrorSchema "User already exists"
// @Router /api/auth/register [post]
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
