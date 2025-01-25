package service

import (
	"context"
	"ridhoandhika/backend-api/domain"
	"ridhoandhika/backend-api/dto"
	"ridhoandhika/backend-api/internal/util"

	"github.com/google/uuid"
)

type skillService struct {
	skillRepository domain.SkillRepository
}

func Skill(skillRepository domain.SkillRepository) domain.SkillService {
	return &skillService{
		skillRepository: skillRepository,
	}
}

func (s skillService) FindByUserId(ctx context.Context, userId string) (dto.BaseResp, error) {
	parsedUserID, err := uuid.Parse(userId)
	if err != nil {
		return util.ErrorResponse("400", "Invalid UUID format", "Invalid UUID format"), nil
	}

	data, err := s.skillRepository.FindByUserId(ctx, parsedUserID)
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

	var response dto.SkillsResp
	var skillList []dto.Skill
	// Convert data ke WorkExperience
	for _, skill := range data {
		// Tambahkan objek Education ke dalam educationList
		skillList = append(skillList, dto.Skill{
			SkillID: skill.SkillID,
			Name:    skill.Name,
			Level:   string(skill.Level),
		})
	}

	response = dto.SkillsResp{
		Skills: skillList,
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

func (s skillService) Insert(ctx context.Context, req dto.InsertSkillReq) (dto.BaseResp, error) {
	_, err := s.skillRepository.Insert(ctx, req)
	if err != nil {
		return util.ErrorResponse("400", "Permintaan Tidak Valid", "Bad Request"), err
	}

	return util.ErrorResponse("200", "Sukses", "Success"), nil
}

func (s skillService) Update(ctx context.Context, skillId string, req dto.UpdateSkillReq) (dto.BaseResp, error) {
	parsedSkillId, err := uuid.Parse(skillId)
	if err != nil {
		return util.ErrorResponse("400", "Permintaan tidak valid", "Bad request"), nil
	}

	_, err = s.skillRepository.Update(ctx, parsedSkillId, req)
	if err != nil {
		return util.ErrorResponse("400", "Gagal", "Failed"), nil
	}

	return util.ErrorResponse("200", "Sukses", "Success"), nil
}
