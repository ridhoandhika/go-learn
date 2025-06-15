package repository

import (
	"context"
	"ridhoandhika/backend-api/domain"
	"ridhoandhika/backend-api/dto"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type certificationRepository struct {
	db *gorm.DB
}

func Certification(con *gorm.DB) domain.CertificationRepository {
	return &certificationRepository{
		db: con,
	}
}

func (r certificationRepository) FindByUserId(ctx context.Context, userId uuid.UUID) ([]domain.Certification, error) {
	var certification []domain.Certification
	err := r.db.WithContext(ctx).Where("user_id = ?", userId).Find(&certification).Error
	return certification, err
}

func (r certificationRepository) Insert(ctx context.Context, req dto.InsertCertificationParsedReq) (bool, error) {
	certification := domain.Certification{
		CertificationID: uuid.New(),
		UserID:          req.UserID,
		Name:            req.Name,
		Body:            req.Body,
		CredentialID:    req.CredentialID,
		IssueDate:       req.IssueDate,
		ExpirationDate:  req.ExpirationDate,
	}

	err := r.db.WithContext(ctx).Create(&certification).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r certificationRepository) Update(ctx context.Context, certificationId uuid.UUID, req dto.UpdateCertificationParsedReq) (bool, error) {
	var certification domain.Certification
	err := r.db.WithContext(ctx).Where("certification_id = ?", certificationId).First(&certification).Error
	if err != nil {
		return false, err
	}

	err = r.db.WithContext(ctx).Model(&certification).Updates(domain.Certification{
		Name:           req.Name,
		Body:           req.Body,
		CredentialID:   req.CredentialID,
		IssueDate:      req.IssueDate,
		ExpirationDate: req.ExpirationDate,
	}).Error

	if err != nil {
		return false, err
	}

	return true, nil
}

func (r certificationRepository) Delete(ctx context.Context, certificationId uuid.UUID) (bool, error) {
	err := r.db.WithContext(ctx).
		Where("certification_id = ?", certificationId).
		Delete(&domain.Certification{}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
