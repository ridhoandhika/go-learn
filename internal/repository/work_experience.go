package repository

import (
	"context"
	"ridhoandhika/backend-api/domain"
	"ridhoandhika/backend-api/dto"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type workExperienceRepository struct {
	db *gorm.DB
}

func WorkExperience(con *gorm.DB) domain.WorkExperienceRepository {
	return &workExperienceRepository{
		db: con,
	}
}

func (u workExperienceRepository) FindByUserId(ctx context.Context, userId uuid.UUID) ([]domain.WorkExperience, error) {
	var workExperiences []domain.WorkExperience
	err := u.db.WithContext(ctx).Where("user_id = ?", userId).Find(&workExperiences).Error
	return workExperiences, err
}

func (u workExperienceRepository) Insert(ctx context.Context, req dto.InsertWorkExperienceReq) (bool, error) {
	workExperience := domain.WorkExperience{
		WorkExperienceID: uuid.New(),
		UserID:           req.UserID,
		JobTitle:         req.JobTitle,
		CompanyName:      req.CompanyName,
		StartDate:        req.StartDate,
		EndDate:          req.EndDate,
		JobDescription:   req.JobDescription,
	}

	err := u.db.WithContext(ctx).Create(&workExperience).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func (u workExperienceRepository) Update(ctx context.Context, workExperienceId uuid.UUID, req dto.UpdateWorkExperienceReq) (bool, error) {
	var workExperience domain.WorkExperience
	err := u.db.WithContext(ctx).Where("work_experience_id = ?", workExperienceId).First(&workExperience).Error
	if err != nil {
		return false, err
	}

	err = u.db.WithContext(ctx).Model(&workExperience).Updates(domain.WorkExperience{
		JobTitle:       req.JobTitle,
		CompanyName:    req.CompanyName,
		StartDate:      req.StartDate,
		EndDate:        req.EndDate,
		JobDescription: req.JobDescription,
	}).Error

	if err != nil {
		return false, err
	}

	return true, nil
}
