package repository

import (
	"context"
	"ridhoandhika/backend-api/domain"
	"ridhoandhika/backend-api/dto"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type personalInformationRepository struct {
	db *gorm.DB
}

func PersonalInformation(con *gorm.DB) domain.PersonalInformationRepository {
	return &personalInformationRepository{
		db: con,
	}
}

func (u personalInformationRepository) FindByID(ctx context.Context, id uuid.UUID) (peronalInformation domain.PersonalInformation, err error) {
	err = u.db.WithContext(ctx).Where("personal_info_id = ?", id).First(&peronalInformation).Error
	return
}

func (u personalInformationRepository) FindByUserID(ctx context.Context, userId uuid.UUID) (peronalInformation domain.PersonalInformation, err error) {
	err = u.db.WithContext(ctx).Where("user_id = ?", userId).First(&peronalInformation).Error
	return
}

func (u personalInformationRepository) Insert(ctx context.Context, req dto.InsertPersonalInformationReq) (interface{}, error) {
	personalInfo := domain.PersonalInformation{
		PersonalInfoID: uuid.New(), // Generate UUID baru
		UserID:         req.UserID,
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		PhoneNumber:    req.PhoneNumber,
		Email:          req.Email,
		Address:        req.Address,
		Summary:        req.Summary,
		DateOfBirth:    req.DateOfBirth,
	}

	// Menyisipkan user baru ke dalam tabel
	err := u.db.WithContext(ctx).Create(&personalInfo).Error
	if err != nil {
		return nil, err
	}

	// Kembalikan ID user yang baru saja dimasukkan
	return nil, nil
}
func (u personalInformationRepository) Update(ctx context.Context, personalInfoID uuid.UUID, req dto.UpdatePersonalInformationReq) (bool, error) {
	var personalInformation domain.PersonalInformation
	err := u.db.WithContext(ctx).Where("personal_info_id = ?", personalInfoID).First(&personalInformation).Error
	if err != nil {
		// Jika data tidak ditemukan, kembalikan error
		return false, err
	}

	// Update data PersonalInformation dengan nilai-nilai yang diberikan dalam request
	err = u.db.WithContext(ctx).Model(&personalInformation).Updates(domain.PersonalInformation{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		PhoneNumber: req.PhoneNumber,
		Email:       req.Email,
		Address:     req.Address,
		Summary:     req.Summary,
		DateOfBirth: req.DateOfBirth,
	}).Error

	if err != nil {
		// Jika ada error saat update
		return false, err
	}

	// Jika berhasil, kembalikan data yang sudah diperbarui
	return true, nil
}
