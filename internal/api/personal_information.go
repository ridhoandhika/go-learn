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
	app.Get("personal_information/:id", authMid, handler.FindByID)
	app.Put("personal_information/:PersonalInfoID", authMid, handler.Update)
	app.Post("personal_information", authMid, handler.Insert)
}

// @Summary Get Personal Information
// @Description Retrieve detailed personal information by ID
// @Tags personal-information
// @Accept json
// @Produce json
// @Security BearerAuth  // Menunjukkan endpoint membutuhkan otentikasi Bearer token
// @Param Authorization header string true "Bearer JWT Token"  // Menambahkan parameter header untuk token otentikasi
// @Param id path string true "Personal Information ID"
// @Success 200 {object} dto.BaseResp{outputSchema=dto.PersonalInformationResp} "Personal Information Details"
// @Failure 400 {object} dto.ErrorSchema "Invalid Request"
// @Failure 401 {object} dto.ErrorSchema "Authentication Failed"
// @Failure 404 {object} dto.ErrorSchema "User Not Found"
// @Router /api/personal_information/{id} [get]
func (a personalInformationhApi) FindByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	data, err := a.personalInformationService.FindByIDPeronalInfo(ctx.Context(), id)
	if err != nil {
		return ctx.SendStatus(util.GetHttpStatus(err))
	}

	return ctx.Status(200).JSON(data)
}

// @Summary Insert Personal Information
// @Description Retrieve detailed personal information by ID
// @Tags personal-information
// @Accept json
// @Produce json
// @Security BearerAuth  // Menunjukkan endpoint membutuhkan otentikasi Bearer token
// @Param Authorization header string true "Bearer JWT Token"  // Menambahkan parameter header untuk token otentikasi
// @Param body body dto.InsertPersonalInformationReq true "Body Request" // Mendefinisikan parameter yang dikirimkan dalam request body
// @Success 200 {object} dto.BaseResp "Personal Information Details"
// @Failure 409 {object} dto.ErrorSchema "Conflict"
// @Failure 400 {object} dto.ErrorSchema "Bad Request"
// @Router /api/personal_information [post]
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

// @Summary Update Personal Information
// @Description Retrieve detailed personal information by ID
// @Tags personal-information
// @Accept json
// @Produce json
// @Security BearerAuth  // Menunjukkan endpoint membutuhkan otentikasi Bearer token
// @Param Authorization header string true "Bearer JWT Token"  // Menambahkan parameter header untuk token otentikasi
// @Param PersonalInfoID path string true "Personal Information ID"
// @Param body body dto.UpdatePersonalInformationReq true "Body Request" // Mendefinisikan parameter yang dikirimkan dalam request body
// @Success 200 {object} dto.BaseResp "Personal Information Details"
// @Failure 400 {object} dto.ErrorSchema "Bad Request"
// @Router /api/personal_information/{PersonalInfoID} [put]
func (a personalInformationhApi) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("PersonalInfoID")
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
