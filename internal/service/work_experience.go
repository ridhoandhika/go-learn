package service

import (
	"context"
	"ridhoandhika/backend-api/domain"
	"ridhoandhika/backend-api/dto"
	"ridhoandhika/backend-api/internal/util"

	"github.com/google/uuid"
)

type workExperienceService struct {
	workExperienceRepository domain.WorkExperienceRepository
}

func WorkExperience(workExperienceRepository domain.WorkExperienceRepository) domain.WorkExperienceService {
	return &workExperienceService{
		workExperienceRepository: workExperienceRepository,
	}
}

func (p workExperienceService) Insert(ctx context.Context, req dto.InsertWorkExperienceReq) (dto.BaseResp, error) {
	_, err := p.workExperienceRepository.Insert(ctx, req)
	if err != nil {
		return util.ErrorResponse("400", "Permintaan Tidak Valid", "Bad Request"), err
	}
	// return response
	return util.ErrorResponse("200", "Sukses", "Success"), nil
}

func (p workExperienceService) FindByUserId(ctx context.Context, userId string) (dto.BaseResp, error) {
	parsedUserID, err := uuid.Parse(userId)
	if err != nil {
		return util.ErrorResponse("400", "Invalid UUID format", "Invalid UUID format"), nil
	}

	data, err := p.workExperienceRepository.FindByUserId(ctx, parsedUserID)
	if err != nil {
		return util.ErrorResponse("404", "Pengguna tidak ditemukan", "User not found"), nil
	}

	var response dto.WorkExperiencesResp
	var workExperienceList []dto.WorkExperience
	// Convert data WorkExperience ke WorkExperiencesResp
	for _, work := range data {
		// Tambahkan objek Education ke dalam educationList
		workExperienceList = append(workExperienceList, dto.WorkExperience{
			JobTitle:       work.JobTitle,
			CompanyName:    work.CompanyName,
			StartDate:      work.StartDate,
			EndDate:        work.EndDate,
			JobDescription: work.JobDescription,
		})
	}

	response = dto.WorkExperiencesResp{
		WorkExperience: workExperienceList,
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

func (w workExperienceService) Update(ctx context.Context, workExperienceId string, req dto.UpdateWorkExperienceReq) (dto.BaseResp, error) {
	parsedWorkExperienceId, err := uuid.Parse(workExperienceId)
	if err != nil {
		// Jika sudah ada, return error bahwa data sudah ada
		return util.ErrorResponse("400", "Permintaan tidak valid", "Bad request"), nil
	}

	_, err = w.workExperienceRepository.Update(ctx, parsedWorkExperienceId, req)
	if err != nil {
		return util.ErrorResponse("400", "Gagal", "Failed"), nil
	}

	// return response
	return util.ErrorResponse("200", "Sukses", "Success"), nil
}
