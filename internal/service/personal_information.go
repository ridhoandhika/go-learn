package service

import (
	"context"
	"ridhoandhika/backend-api/domain"
	"ridhoandhika/backend-api/dto"
	"ridhoandhika/backend-api/internal/util"

	"github.com/google/uuid"
)

type personalInformationService struct {
	personaleInformationRepository domain.PersonalInformationRepository
}

func PersonalInformation(personalInformationRepository domain.PersonalInformationRepository) domain.PersonalInformationService {
	return &personalInformationService{
		personaleInformationRepository: personalInformationRepository,
	}
}

func (u personalInformationService) FindByIDPeronalInfo(ctx context.Context, id string) (dto.BaseResp, error) {
	parsedID, err := uuid.Parse(id)
	data, err := u.personaleInformationRepository.FindByID(ctx, parsedID)
	//  u.personalInformationRepository.FindByID(ctx, id)
	if err != nil {
		return util.ErrorResponse("404", "User tidak ditemukan", "User not found"), nil
	}
	if err != nil {
		return util.ErrorResponse("404", "User tidak ditemukan", "User not found"), nil
	}

	return dto.BaseResp{
		ErrorSchema: dto.ErrorSchema{
			ErrorCode: "200",
			ErrorMessage: dto.ErrorMessage{
				Id: "Sukes",
				En: "Success",
			},
		},
		OutputSchema: dto.PersonalInformationResp{
			PersonalInfoID: data.PersonalInfoID,
			UserID:         data.UserID,
			FirstName:      data.FirstName,
			LastName:       data.LastName,
			PhoneNumber:    data.PhoneNumber,
			Email:          data.Email,
			Address:        data.Address,
			Summary:        data.Summary,
			DateOfBirth:    data.DateOfBirth,
		},
	}, nil
}

func (p personalInformationService) Insert(ctx context.Context, req dto.InsertPersonalInformationReq) (dto.BaseResp, error) {
	_, err := p.personaleInformationRepository.FindByUserID(ctx, req.UserID)
	if err == nil {
		// Jika sudah ada, return error bahwa data sudah ada
		return util.ErrorResponse("409", "Konflik", "Conflict"), nil
	}

	_, err = p.personaleInformationRepository.Insert(ctx, req)
	if err != nil {
		return util.ErrorResponse("400", "Permintaan Tidak Valid", "Bad Request"), err
	}
	// return response
	return util.ErrorResponse("200", "Sukses", "Success"), nil
}

func (p personalInformationService) Update(ctx context.Context, personalInfoID uuid.UUID, req dto.UpdatePersonalInformationReq) (dto.BaseResp, error) {
	_, err := p.personaleInformationRepository.Update(ctx, personalInfoID, req)
	if err != nil {
		// Jika sudah ada, return error bahwa data sudah ada
		return util.ErrorResponse("400", "Permintaan Tidak Valid", "Bad Request"), nil
	}

	// return response
	return util.ErrorResponse("200", "Sukses", "Success"), nil
}
