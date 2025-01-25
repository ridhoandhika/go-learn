package repository

import (
	"context"
	"ridhoandhika/backend-api/domain"
	"ridhoandhika/backend-api/dto"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type skillRepository struct {
	db *gorm.DB
}

func Skill(con *gorm.DB) domain.SkillRepository {
	return &skillRepository{
		db: con,
	}
}

func (r skillRepository) FindByUserId(ctx context.Context, userId uuid.UUID) ([]domain.Skill, error) {
	var skills []domain.Skill
	err := r.db.WithContext(ctx).Where("user_id = ?", userId).Find(&skills).Error
	return skills, err
}

func (r skillRepository) Insert(ctx context.Context, req dto.InsertSkillReq) (bool, error) {
	skill := domain.Skill{
		SkillID: uuid.New(),
		UserID:  req.UserID,
		Name:    req.Name,
		Level:   domain.Level(req.Level),
	}

	err := r.db.WithContext(ctx).Create(&skill).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r skillRepository) Update(ctx context.Context, skillId uuid.UUID, req dto.UpdateSkillReq) (bool, error) {
	var skill domain.Skill
	err := r.db.WithContext(ctx).Where("skill_id = ?", skillId).First(&skill).Error
	if err != nil {
		return false, err
	}

	err = r.db.WithContext(ctx).Model(&skill).Updates(domain.Skill{
		Name:  req.Name,
		Level: domain.Level(req.Level),
	}).Error

	if err != nil {
		return false, err
	}

	return true, nil
}
