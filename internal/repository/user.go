package repository

import (
	"context"
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

	// user := domain.User{
	// 	Username: req.Username,
	// 	Password: req.Password,
	// 	Phone:    req.Phone,
	// 	Fullname: req.Fullname,
	// }

	user := domain.User{
		ID:       uuid.New(), // Generate UUID baru
		Username: req.Username,
		Password: req.Password,
		Phone:    req.Phone,
		Fullname: req.Fullname,
	}

	// Menyisipkan user baru ke dalam tabel
	err := u.db.WithContext(ctx).Create(&user).Error
	if err != nil {
		return nil, err
	}

	// Kembalikan ID user yang baru saja dimasukkan
	return nil, nil
}
