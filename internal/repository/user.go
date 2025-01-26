package repository

import (
	"context"
	"errors"
	"ridhoandhika/backend-api/domain"
	"ridhoandhika/backend-api/dto"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func User(con *gorm.DB) domain.UserRepository {
	return &userRepository{
		db: con,
	}
}

func (u userRepository) FindByID(ctx context.Context, id uuid.UUID) (user domain.User, err error) {
	err = u.db.WithContext(ctx).Where("id = ?", id).First(&user).Error
	return
}

func (u userRepository) FindByUsername(ctx context.Context, username string) (user domain.User, err error) {
	err = u.db.WithContext(ctx).Where("username = ?", username).First(&user).Error
	return
}

func (u userRepository) InsertUser(ctx context.Context, req dto.UserRegisterReq) (interface{}, error) {

	user := domain.User{
		ID:       uuid.New(),
		Username: req.Username,
		Password: req.Password,
		Phone:    req.Phone,
		Fullname: req.Fullname,
	}

	err := u.db.WithContext(ctx).Create(&user).Error
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (u userRepository) UserWithEducation(ctx context.Context, userID uuid.UUID) (*domain.User, error) {
	var user domain.User
	// Cari user berdasarkan ID dan preload hanya data Education
	if err := u.db.Preload("Education").First(&user, "id = ?", userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

func (u userRepository) UserWithWorkExperience(ctx context.Context, userID uuid.UUID) (*domain.User, error) {
	var user domain.User
	// Cari user berdasarkan ID dan preload hanya data Education
	if err := u.db.Preload("WorkExperience").First(&user, "id = ?", userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

func (u userRepository) UserWithSkill(ctx context.Context, userID uuid.UUID) (*domain.User, error) {
	var user domain.User
	// Cari user berdasarkan ID dan preload hanya data Education
	if err := u.db.Preload("Skill").First(&user, "id = ?", userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

func (u userRepository) UserWithPersonalInformation(ctx context.Context, userID uuid.UUID) (*domain.User, error) {
	var user domain.User
	// Cari user berdasarkan ID dan preload hanya data Education
	if err := u.db.Preload("PersonalInfo").First(&user, "id = ?", userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

func (u userRepository) UserWithCertification(ctx context.Context, userID uuid.UUID) (*domain.User, error) {
	var user domain.User
	// Cari user berdasarkan ID dan preload hanya data Education
	if err := u.db.Preload("Certification").First(&user, "id = ?", userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}
