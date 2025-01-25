package repository

import (
	"context"
	"ridhoandhika/backend-api/domain"
	"ridhoandhika/backend-api/dto"
	"ridhoandhika/backend-api/internal/util"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const formatDate = "2006-01-02"

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

func (r certificationRepository) Insert(ctx context.Context, req dto.InsertCertificationReq) (bool, error) {

	// Parsing IssueDate menggunakan utilitas ParseDate
	issueDate, err := util.ParseDate(req.IssueDate, formatDate)
	if err != nil {
		return false, err
	}

	// Parsing ExpirationDate menggunakan utilitas ParseDate
	expirationDate, err := util.ParseDate(req.ExpirationDate, formatDate)
	if err != nil {
		return false, err
	}

	certification := domain.Certification{
		CertificationID: uuid.New(),
		UserID:          req.UserID,
		Name:            req.Name,
		Body:            req.Body,
		CredentialID:    req.CredentialID,
		IssueDate:       issueDate,
		ExpirationDate:  expirationDate,
	}

	err = r.db.WithContext(ctx).Create(&certification).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r certificationRepository) Update(ctx context.Context, certificationId uuid.UUID, req dto.UpdateCertificationReq) (bool, error) {
	var certification domain.Certification
	err := r.db.WithContext(ctx).Where("certification_id = ?", certificationId).First(&certification).Error
	if err != nil {
		return false, err
	}

	// Parsing IssueDate menggunakan utilitas ParseDate
	issueDate, err := util.ParseDate(req.IssueDate, formatDate)
	if err != nil {
		return false, err
	}

	// Parsing ExpirationDate menggunakan utilitas ParseDate
	expirationDate, err := util.ParseDate(req.ExpirationDate, formatDate)
	if err != nil {
		return false, err
	}

	err = r.db.WithContext(ctx).Model(&certification).Updates(domain.Certification{
		CertificationID: uuid.New(),
		Name:            req.Name,
		Body:            req.Body,
		IssueDate:       issueDate,
		ExpirationDate:  expirationDate,
	}).Error

	if err != nil {
		return false, err
	}

	return true, nil
}
