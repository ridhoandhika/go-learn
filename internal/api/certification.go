package api

import (
	"fmt"
	"ridhoandhika/backend-api/domain"
	"ridhoandhika/backend-api/dto"
	"ridhoandhika/backend-api/internal/util"

	"github.com/gofiber/fiber/v2"
)

type certificationApi struct {
	certificationService domain.CertificationService
}

func Certification(app *fiber.Group, certificationService domain.CertificationService, authMid fiber.Handler) {
	handler := certificationApi{
		certificationService: certificationService,
	}
	app.Get("certification/:userId", authMid, handler.FindByUserId)
	app.Post("certification", authMid, handler.Insert)
	app.Put("certification/:id", authMid, handler.Update)
}

// @Security BearerAuth
// @Summary Get Certification
// @Tags certification
// @Accept json
// @Produce json
// @Param userId path string true "User ID"
// @Success 200 {object} dto.BaseResp{output_schema=dto.CertificationResp{certifications=[]dto.Certification{}}} "Certification details successfully retrieved"
// @Failure 400 {object} dto.ErrorSchema "Bad Request"
// @Router /api/certification/{userId} [get]
func (a certificationApi) FindByUserId(ctx *fiber.Ctx) error {
	id := ctx.Params("userId")
	data, err := a.certificationService.FindByUserId(ctx.Context(), id)
	if err != nil {
		// Log error for debugging
		fmt.Printf("Error: %v\n", err)
		return ctx.SendStatus(util.GetHttpStatus(err))
	}

	return ctx.Status(200).JSON(data)
}

// @Security BearerAuth
// @Summary Insert Certification
// @Tags certification
// @Accept json
// @Produce json
// @Param body body dto.InsertCertificationReq true "Body Request"
// @Success 200 {object} dto.BaseResp "Certification Details"
// @Failure 400 {object} dto.ErrorSchema "Bad Request"
// @Router /api/certification [post]
func (a certificationApi) Insert(ctx *fiber.Ctx) error {
	var req dto.InsertCertificationReq
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(200).JSON(util.ErrorResponse("400", "Permintaan Tidak Valid", "Bad Request"))
	}

	data, err := a.certificationService.Insert(ctx.Context(), req)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return ctx.SendStatus(util.GetHttpStatus(err))
	}

	return ctx.Status(200).JSON(data)
}

// @Security BearerAuth
// @Summary Update Certification
// @Tags certification
// @Accept json
// @Produce json
// @Param id path string true "Certification ID"
// @Param body body dto.UpdateCertificationReq true "Body Request"
// @Success 200 {object} dto.BaseResp "Certification resp"
// @Failure 400 {object} dto.ErrorSchema "Bad Request"
// @Router /api/certification/{id} [put]
func (a certificationApi) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var req dto.UpdateCertificationReq
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(200).JSON(util.ErrorResponse("400", "Permintaan Tidak Valid", "Bad Request"))
	}

	data, err := a.certificationService.Update(ctx.Context(), id, req)

	if err != nil {
		return ctx.SendStatus(util.GetHttpStatus(err))
	}

	return ctx.Status(200).JSON(data)
}
