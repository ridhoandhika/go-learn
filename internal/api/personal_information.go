package api

import (
	"ridhoandhika/backend-api/domain"
	"ridhoandhika/backend-api/dto"
	"ridhoandhika/backend-api/internal/util"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type personalInformationhApi struct {
	personalInformationService domain.PersonalInformationService
}

func PersonalInformation(app *fiber.Group, personalInformationService domain.PersonalInformationService, authMid fiber.Handler) {
	handler := personalInformationhApi{
		personalInformationService: personalInformationService,
	}
	app.Get("personal-information/:id", authMid, handler.FindByID)
	app.Put("personal-information/:id", authMid, handler.Update)
	app.Post("personal-information", authMid, handler.Insert)
}

// @Security BearerAuth
// @Summary Get Personal Information
// @Tags personal-information
// @Accept json
// @Produce json
// @Param id path string true "Personal Information ID"
// @Success 200 {object} dto.BaseResp{output_schema=dto.PersonalInformationResp} "Personal Information Details"
// @Failure 400 {object} dto.ErrorSchema "Invalid Request"
// @Failure 401 {object} dto.ErrorSchema "Authentication Failed"
// @Failure 404 {object} dto.ErrorSchema "User Not Found"
// @Router /api/personal-information/{id} [get]
func (a personalInformationhApi) FindByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	data, err := a.personalInformationService.FindByIDPeronalInfo(ctx.Context(), id)
	if err != nil {
		return ctx.SendStatus(util.GetHttpStatus(err))
	}

	return ctx.Status(200).JSON(data)
}

// @Security BearerAuth
// @Summary Insert Personal Information
// @Tags personal-information
// @Accept json
// @Produce json
// @Param body body dto.InsertPersonalInformationReq true "Body Request"
// @Success 200 {object} dto.BaseResp "Personal Information Details"
// @Failure 409 {object} dto.ErrorSchema "Conflict"
// @Failure 400 {object} dto.ErrorSchema "Bad Request"
// @Router /api/personal-information [post]
func (a personalInformationhApi) Insert(ctx *fiber.Ctx) error {
	var req dto.InsertPersonalInformationReq
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(200).JSON(util.ErrorResponse("400", "Permintaan Tidak Valid", "Bad Request"))
	}

	data, err := a.personalInformationService.Insert(ctx.Context(), req)
	if err != nil {
		return ctx.Status(200).JSON(util.ErrorResponse("400", "Permintaan Tidak Valid", "Bad Request"))
	}

	return ctx.Status(200).JSON(data)
}

// @Security BearerAuth
// @Summary Update Personal Information
// @Tags personal-information
// @Accept json
// @Produce json
// @Param id path string true "Personal Information ID"
// @Param body body dto.UpdatePersonalInformationReq true "Body Request"
// @Success 200 {object} dto.BaseResp "Personal Information Details"
// @Failure 400 {object} dto.ErrorSchema "Bad Request"
// @Router /api/personal-information/{id} [put]
func (a personalInformationhApi) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	// fmt.Printf("personalInfoID " + ctx.Params("id"))
	var req dto.UpdatePersonalInformationReq
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(200).JSON(util.ErrorResponse("400", "Permintaan Tidak Valid", "Bad Request"))
	}

	personalInfoID, _ := uuid.Parse(id)

	data, err := a.personalInformationService.Update(ctx.Context(), personalInfoID, req)

	if err != nil {
		return ctx.SendStatus(util.GetHttpStatus(err))
	}

	return ctx.Status(200).JSON(data)
}
