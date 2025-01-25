package service

import (
	"context"
	"ridhoandhika/backend-api/domain"
	"ridhoandhika/backend-api/dto"
	"ridhoandhika/backend-api/internal/util"

	"github.com/google/uuid"
)

type educationService struct {
	educationRepository domain.EducationRepository
}

func Education(educationRepository domain.EducationRepository) domain.EducationService {
	return &educationService{
		educationRepository: educationRepository,
	}
}

func (p educationService) FindByUserId(ctx context.Context, userId string) (dto.BaseResp, error) {

	parsedUserID, err := uuid.Parse(userId)
	if err != nil {
		return util.ErrorResponse("400", "Invalid UUID format", "Invalid UUID format"), nil
	}

	data, err := p.educationRepository.FindByUserId(ctx, parsedUserID)

	if err != nil || len(data) == 0 {
		return dto.BaseResp{
			ErrorSchema: dto.ErrorSchema{
				ErrorCode: "200",
				ErrorMessage: dto.ErrorMessage{
					Id: "Sukses",
					En: "Success",
				},
			},
			OutputSchema: dto.EducationResp{
				Education: []dto.Education{},
			},
		}, nil
	}

	var response dto.EducationResp
	var educationList []dto.Education
	// Convert data ke Education
	for _, education := range data {
		// Tambahkan objek Education ke dalam educationList
		educationList = append(educationList, dto.Education{
			EducationID:  education.EducationID,
			SchoolName:   education.SchoolName,
			Degree:       education.Degree,
			FieldOfStudy: education.FieldOfStudy,
			Description:  education.Description,
			StartDate:    education.StartDate,
			EndDate:      education.EndDate,
		})
	}

	response = dto.EducationResp{
		Education: educationList,
	}

	return dto.BaseResp{
		ErrorSchema: dto.ErrorSchema{
			ErrorCode: "200",
			ErrorMessage: dto.ErrorMessage{
				Id: "Success",
				En: "Success",
			},
		},
		OutputSchema: response,
	}, nil
}

func (p educationService) Insert(ctx context.Context, req dto.InsertEducationReq) (dto.BaseResp, error) {
	_, err := p.educationRepository.Insert(ctx, req)
	if err != nil {
		return util.ErrorResponse("400", "Permintaan Tidak Valid", "Bad Request"), err
	}

	return util.ErrorResponse("200", "Sukses", "Success"), nil
}

func (w educationService) Update(ctx context.Context, educationId string, req dto.UpdateEducationReq) (dto.BaseResp, error) {
	parsedEducationId, err := uuid.Parse(educationId)
	if err != nil {
		return util.ErrorResponse("400", "Permintaan tidak valid", "Bad request"), nil
	}

	_, err = w.educationRepository.Update(ctx, parsedEducationId, req)
	if err != nil {
		return util.ErrorResponse("400", "Gagal", "Failed"), nil
	}

	return util.ErrorResponse("200", "Sukses", "Success"), nil
}
