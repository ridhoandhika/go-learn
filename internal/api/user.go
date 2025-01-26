package api

import (
	"ridhoandhika/backend-api/domain"
	"ridhoandhika/backend-api/internal/util"

	"github.com/gofiber/fiber/v2"
)

type userApi struct {
	userService domain.UserService
}

func User(app *fiber.Group, userService domain.UserService, authMid fiber.Handler) {
	handler := userApi{
		userService: userService,
	}
	app.Get("user/:id", authMid, handler.GetUser)
	app.Get("user/:id/education", authMid, handler.GetEducation)
	app.Get("user/:id/personal-information", authMid, handler.GetPersonalInformation)
	app.Get("user/:id/work-experience", authMid, handler.GetWorkExperience)
	app.Get("user/:id/skill", authMid, handler.GetSkill)
	app.Get("user/:id/certification", authMid, handler.GetCertification)
}

// @Security BearerAuth
// @Summary Get user by id
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "user id"
// @Success 200 {object} dto.BaseResp "User Details"
// @Failure 400 {object} dto.ErrorSchema "Bad Request"
// @Router /api/user/{id} [get]
func (a userApi) GetUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	data, err := a.userService.GetUser(ctx.Context(), id)
	if err != nil {
		return ctx.Status(200).JSON(util.ErrorResponse("400", "Permintaan Tidak Valid", "Bad Request"))
	}

	return ctx.Status(200).JSON(data)
}

// @Security BearerAuth
// @Summary Get Education by user id
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "user id"
// @Success 200 {object} dto.BaseResp{output_schema=dto.UserEducationResp{educations=[]dto.Education{}}} "Education Details"
// @Failure 400 {object} dto.ErrorSchema "Bad Request"
// @Router /api/user/{id}/education [get]
func (a userApi) GetEducation(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	data, err := a.userService.GetEducation(ctx.Context(), id)
	if err != nil {
		return ctx.Status(200).JSON(util.ErrorResponse("400", "Permintaan Tidak Valid", "Bad Request"))
	}

	return ctx.Status(200).JSON(data)
}

// @Security BearerAuth
// @Summary Get personal information by user id
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "user id"
// @Success 200 {object} dto.BaseResp "Personal Information Details"
// @Failure 400 {object} dto.ErrorSchema "Bad Request"
// @Router /api/user/{id}/personal-information [get]
func (a userApi) GetPersonalInformation(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	data, err := a.userService.GetPersonalInformation(ctx.Context(), id)
	if err != nil {
		return ctx.Status(200).JSON(util.ErrorResponse("400", "Permintaan Tidak Valid", "Bad Request"))
	}

	return ctx.Status(200).JSON(data)
}

// @Security BearerAuth
// @Summary Get work experience by user id
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "user id"
// @Success 200 {object} dto.BaseResp{output_schema=dto.UserWorkExperienceResp{work_experiences=[]dto.WorkExperience{}}} "Work Experience Details"
// @Failure 400 {object} dto.ErrorSchema "Bad Request"
// @Router /api/user/{id}/work-experience [get]
func (a userApi) GetWorkExperience(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	data, err := a.userService.GetWorkExperience(ctx.Context(), id)
	if err != nil {
		return ctx.Status(200).JSON(util.ErrorResponse("400", "Permintaan Tidak Valid", "Bad Request"))
	}

	return ctx.Status(200).JSON(data)
}

// @Security BearerAuth
// @Summary Get skill by user id
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "user id"
// @Success 200 {object} dto.BaseResp{output_schema=dto.UserSkillResp{skills=[]dto.Skill{}}} "Skill Details"
// @Failure 400 {object} dto.ErrorSchema "Bad Request"
// @Router /api/user/{id}/skill [get]
func (a userApi) GetSkill(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	data, err := a.userService.GetEducation(ctx.Context(), id)
	if err != nil {
		return ctx.Status(200).JSON(util.ErrorResponse("400", "Permintaan Tidak Valid", "Bad Request"))
	}

	return ctx.Status(200).JSON(data)
}

// @Security BearerAuth
// @Summary Get certification by user id
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "user id"
// @Success 200 {object} dto.BaseResp "Certification Details"
// @Failure 400 {object} dto.ErrorSchema "Bad Request"
// @Router /api/user/{id}/certification [get]
func (a userApi) GetCertification(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	data, err := a.userService.GetCertification(ctx.Context(), id)
	if err != nil {
		return ctx.Status(200).JSON(util.ErrorResponse("400", "Permintaan Tidak Valid", "Bad Request"))
	}

	return ctx.Status(200).JSON(data)
}
