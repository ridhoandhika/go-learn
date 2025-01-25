package api

import (
	"fmt"
	"ridhoandhika/backend-api/domain"
	"ridhoandhika/backend-api/dto"
	"ridhoandhika/backend-api/internal/util"

	"github.com/gofiber/fiber/v2"
)

type educationApi struct {
	educationService domain.EducationService
}

func Education(app *fiber.Group, educationService domain.EducationService, authMid fiber.Handler) {
	handler := educationApi{
		educationService: educationService,
	}
	app.Get("education/:userId", authMid, handler.FindByUserId)
	app.Put("education/:id", authMid, handler.Update)
	app.Post("education", authMid, handler.Insert)
}

// @Security BearerAuth
// @Summary Get Education
// @Tags education
// @Accept json
// @Produce json
// @Param userId path string true "User ID"
// @Success 200 {object} dto.BaseResp{output_schema=dto.EducationResp{education=[]dto.Education{}}} "Education details successfully retrieved"
// @Failure 400 {object} dto.ErrorSchema "Bad Request"
// @Router /api/education/{userId} [get]
func (a educationApi) FindByUserId(ctx *fiber.Ctx) error {
	id := ctx.Params("userId")
	data, err := a.educationService.FindByUserId(ctx.Context(), id)
	if err != nil {
		// Log error for debugging
		fmt.Printf("Error: %v\n", err)
		return ctx.SendStatus(util.GetHttpStatus(err))
	}

	return ctx.Status(200).JSON(data)
}

// @Security BearerAuth
// @Summary Insert Education
// @Tags education
// @Accept json
// @Produce json
// @Param body body dto.InsertEducationReq true "Body Request"
// @Success 200 {object} dto.BaseResp "Education Details"
// @Failure 400 {object} dto.ErrorSchema "Bad Request"
// @Router /api/education [post]
func (a educationApi) Insert(ctx *fiber.Ctx) error {
	var req dto.InsertEducationReq
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(200).JSON(util.ErrorResponse("400", "Permintaan Tidak Valid", "Bad Request"))
	}

	data, err := a.educationService.Insert(ctx.Context(), req)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return ctx.SendStatus(util.GetHttpStatus(err))
	}

	return ctx.Status(200).JSON(data)
}

// @Security BearerAuth
// @Summary Update Education
// @Tags education
// @Accept json
// @Produce json
// @Param id path string true "Education ID"
// @Param body body dto.UpdateEducationReq true "Body Request"
// @Success 200 {object} dto.BaseResp "Work Experience resp"
// @Failure 400 {object} dto.ErrorSchema "Bad Request"
// @Router /api/education/{id} [put]
func (a educationApi) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var req dto.UpdateEducationReq
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(200).JSON(util.ErrorResponse("400", "Permintaan Tidak Valid", "Bad Request"))
	}

	data, err := a.educationService.Update(ctx.Context(), id, req)

	if err != nil {
		return ctx.SendStatus(util.GetHttpStatus(err))
	}

	return ctx.Status(200).JSON(data)
}
