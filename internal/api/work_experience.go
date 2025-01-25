package api

import (
	"ridhoandhika/backend-api/domain"
	"ridhoandhika/backend-api/dto"
	"ridhoandhika/backend-api/internal/util"

	"github.com/gofiber/fiber/v2"
)

type workExperienceApi struct {
	workExperienceService domain.WorkExperienceService
}

func WorkExperience(app *fiber.Group, workExperienceService domain.WorkExperienceService, authMid fiber.Handler) {
	handler := workExperienceApi{
		workExperienceService: workExperienceService,
	}
	app.Get("work-experience/:userId", authMid, handler.FindByUserId)
	app.Put("work-experience/:id", authMid, handler.Update)
	app.Post("work-experience", authMid, handler.Insert)
}

// @Summary Insert Work Experience
// @Description Create Work Experience
// @Tags work-experience
// @Accept json
// @Produce json
// @Security BearerAuth  // Menunjukkan endpoint membutuhkan otentikasi Bearer token
// @Param Authorization header string true "Bearer JWT Token"  // Menambahkan parameter header untuk token otentikasi
// @Param body body dto.InsertWorkExperienceReq true "Body Request" // Mendefinisikan parameter yang dikirimkan dalam request body
// @Success 200 {object} dto.BaseResp "Personal Information Details"
// @Failure 409 {object} dto.ErrorSchema "Conflict"
// @Failure 400 {object} dto.ErrorSchema "Bad Request"
// @Router /api/work-experience [post]
func (a workExperienceApi) Insert(ctx *fiber.Ctx) error {
	var req dto.InsertWorkExperienceReq
	// fmt.Printf("requestd " + ctx.BodyParser(&req))
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(200).JSON(util.ErrorResponse("400", "Permintaan Tidak Valid", "Bad Request"))
	}

	data, err := a.workExperienceService.Insert(ctx.Context(), req)
	if err != nil {
		return ctx.Status(200).JSON(util.ErrorResponse("400", "Permintaan Tidak Valid", "Bad Request"))
	}

	return ctx.Status(200).JSON(data)
}

// @Summary Get Work Experience
// @Description Get Work Experience by User Id
// @Tags work-experience
// @Accept json
// @Produce json
// @Security BearerAuth  // Menunjukkan endpoint membutuhkan otentikasi Bearer token
// @Param Authorization header string true "Bearer JWT Token"  // Menambahkan parameter header untuk token otentikasi
// @Param userId path string true "User ID"
// @Success 200 {object} dto.BaseResp{output_schema=dto.WorkExperiencesResp{work_experience=[]dto.WorkExperience{}}} "Working Experiences"
// @Failure 409 {object} dto.ErrorSchema "Conflict"
// @Failure 400 {object} dto.ErrorSchema "Bad Request"
// @Router /api/work-experience/{userId} [get]
func (a workExperienceApi) FindByUserId(ctx *fiber.Ctx) error {
	id := ctx.Params("userId")
	data, err := a.workExperienceService.FindByUserId(ctx.Context(), id)
	if err != nil {
		return ctx.Status(200).JSON(util.ErrorResponse("400", "Permintaan Tidak Valid", "Bad Request"))
	}

	return ctx.Status(200).JSON(data)
}

// @Summary Update Work Experience
// @Description Retrieve detailed work experience by ID
// @Tags work-experience
// @Accept json
// @Produce json
// @Security BearerAuth  // Menunjukkan endpoint membutuhkan otentikasi Bearer token
// @Param Authorization header string true "Bearer JWT Token"  // Menambahkan parameter header untuk token otentikasi
// @Param id path string true "Work Experience ID"
// @Param body body dto.UpdateWorkExperienceReq true "Body Request" // Mendefinisikan parameter yang dikirimkan dalam request body
// @Success 200 {object} dto.BaseResp "Work Experience resp"
// @Failure 400 {object} dto.ErrorSchema "Bad Request"
// @Router /api/work-experience/{id} [put]
func (a workExperienceApi) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var req dto.UpdateWorkExperienceReq
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(200).JSON(util.ErrorResponse("400", "Permintaan Tidak Valid", "Bad Request"))
	}

	data, err := a.workExperienceService.Update(ctx.Context(), id, req)

	if err != nil {
		return ctx.SendStatus(util.GetHttpStatus(err))
	}

	return ctx.Status(200).JSON(data)
}
