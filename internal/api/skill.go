package api

import (
	"fmt"
	"ridhoandhika/backend-api/domain"
	"ridhoandhika/backend-api/dto"
	"ridhoandhika/backend-api/internal/util"

	"github.com/gofiber/fiber/v2"
)

type skillApi struct {
	skillService domain.SkillService
}

func Skill(app *fiber.Group, skillService domain.SkillService, authMid fiber.Handler) {
	handler := skillApi{
		skillService: skillService,
	}
	app.Get("skill/:userId", authMid, handler.FindByUserId)
	app.Post("skill", authMid, handler.Insert)
	app.Put("skill/:id", authMid, handler.Update)
}

// @Security BearerAuth
// @Summary Get Skill
// @Tags skill
// @Accept json
// @Produce json
// @Param userId path string true "User ID"
// @Success 200 {object} dto.BaseResp{output_schema=dto.SkillsResp{skills=[]dto.Skill{}}} "Skill details successfully retrieved"
// @Failure 400 {object} dto.ErrorSchema "Bad Request"
// @Router /api/skill/{userId} [get]
func (a skillApi) FindByUserId(ctx *fiber.Ctx) error {
	id := ctx.Params("userId")
	data, err := a.skillService.FindByUserId(ctx.Context(), id)
	if err != nil {
		// Log error for debugging
		fmt.Printf("Error: %v\n", err)
		return ctx.SendStatus(util.GetHttpStatus(err))
	}

	return ctx.Status(200).JSON(data)
}

// @Security BearerAuth
// @Summary Insert Skill
// @Tags skill
// @Accept json
// @Produce json
// @Param body body dto.InsertSkillReq true "Body Request"
// @Success 200 {object} dto.BaseResp "Skill Details"
// @Failure 400 {object} dto.ErrorSchema "Bad Request"
// @Router /api/skill [post]
func (a skillApi) Insert(ctx *fiber.Ctx) error {
	var req dto.InsertSkillReq
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(200).JSON(util.ErrorResponse("400", "Permintaan Tidak Valid", "Bad Request"))
	}

	data, err := a.skillService.Insert(ctx.Context(), req)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return ctx.SendStatus(util.GetHttpStatus(err))
	}

	return ctx.Status(200).JSON(data)
}

// @Security BearerAuth
// @Summary Update Skill
// @Tags skill
// @Accept json
// @Produce json
// @Param id path string true "Skill ID"
// @Param body body dto.UpdateSkillReq true "Body Request"
// @Success 200 {object} dto.BaseResp "Skill resp"
// @Failure 400 {object} dto.ErrorSchema "Bad Request"
// @Router /api/skill/{id} [put]
func (a skillApi) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var req dto.UpdateSkillReq
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(200).JSON(util.ErrorResponse("400", "Permintaan Tidak Valid", "Bad Request"))
	}

	data, err := a.skillService.Update(ctx.Context(), id, req)

	if err != nil {
		return ctx.SendStatus(util.GetHttpStatus(err))
	}

	return ctx.Status(200).JSON(data)
}
