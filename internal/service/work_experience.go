package service

import (
	"context"
	"ridhoandhika/backend-api/domain"
	"ridhoandhika/backend-api/dto"
	"ridhoandhika/backend-api/internal/util"
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
	// if err == nil {
	// 	// Jika sudah ada, return error bahwa data sudah ada
	// 	return util.ErrorResponse("409", "Konflik", "Conflict"), nil
	// }

	// _, err = p.personaleInformationRepository.Insert(ctx, req)
	if err != nil {
		return util.ErrorResponse("400", "Permintaan Tidak Valid", "Bad Request"), err
	}
	// return response
	return util.ErrorResponse("200", "Sukses", "Success"), nil
}
