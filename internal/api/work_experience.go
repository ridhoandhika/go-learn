package api

import (
	"fmt"
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
	// app.Get("personal_information/:id", authMid, handler.FindByID)
	// app.Put("personal_information/:PersonalInfoID", authMid, handler.Update)
	app.Post("work_experience", authMid, handler.Insert)
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
// @Router /api/work_experience [post]
func (a workExperienceApi) Insert(ctx *fiber.Ctx) error {
	var req dto.InsertWorkExperienceReq
	// fmt.Printf("requestd " + ctx.BodyParser(&req))
	if err := ctx.BodyParser(&req); err != nil {
		fmt.Printf("requestd " + err.Error())
		return ctx.Status(200).JSON(util.ErrorResponse("400", "Permintaan Tidak Valid", "Bad Request"))
	}

	data, err := a.workExperienceService.Insert(ctx.Context(), req)
	if err != nil {
		return ctx.Status(200).JSON(util.ErrorResponse("400", "Permintaan Tidak Valid", "Bad Request"))
	}

	return ctx.Status(200).JSON(data)
}
