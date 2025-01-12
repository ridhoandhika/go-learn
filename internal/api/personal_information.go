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
	app.Put("personal_information/:personalInfoID", authMid, handler.Update)
	app.Post("personal_information", authMid, handler.Insert)
}

func (a personalInformationhApi) FindByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	data, err := a.personalInformationService.FindByIDPeronalInfo(ctx.Context(), id)
	if err != nil {
		return ctx.SendStatus(util.GetHttpStatus(err))
	}

	return ctx.Status(200).JSON(data)
}

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

func (a personalInformationhApi) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("personalInfoID")
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
