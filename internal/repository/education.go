package repository

import (
	"context"
	"ridhoandhika/backend-api/domain"
	"ridhoandhika/backend-api/dto"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type educationRepository struct {
	db *gorm.DB
}

func Education(con *gorm.DB) domain.EducationRepository {
	return &educationRepository{
		db: con,
	}
}

func (u educationRepository) FindByUserId(ctx context.Context, userId uuid.UUID) ([]domain.Education, error) {
	var educations []domain.Education
	err := u.db.WithContext(ctx).Where("user_id = ?", userId).Find(&educations).Error
	if err != nil {
		return nil, err
	}
	return educations, nil
}

func (u educationRepository) Insert(ctx context.Context, req dto.InsertEducationReq) (bool, error) {
	education := domain.Education{
		EducationID:  uuid.New(),
		UserID:       req.UserID,
		Degree:       req.Degree,
		SchoolName:   req.SchoolName,
		FieldOfStudy: req.FieldOfStudy,
		Description:  req.Description,
		StartDate:    req.StartDate,
		EndDate:      req.EndDate,
	}

	err := u.db.WithContext(ctx).Create(&education).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func (u educationRepository) Update(ctx context.Context, educationId uuid.UUID, req dto.UpdateEducationReq) (bool, error) {
	var education domain.Education
	err := u.db.WithContext(ctx).Where("education_id = ?", educationId).First(&education).Error
	if err != nil {
		return false, err
	}

	err = u.db.WithContext(ctx).Model(&education).Updates(domain.Education{
		Degree:       req.Degree,
		SchoolName:   req.SchoolName,
		FieldOfStudy: req.FieldOfStudy,
		Description:  req.Description,
		StartDate:    req.StartDate,
		EndDate:      req.EndDate,
	}).Error

	if err != nil {
		return false, err
	}

	return true, nil
}
