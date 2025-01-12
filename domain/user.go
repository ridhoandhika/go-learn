package domain

import (
	"context"
	"ridhoandhika/backend-api/dto"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`                   // UUID sebagai primary key
	Username  string    `gorm:"type:varchar(100);uniqueIndex;not null"` // Kolom Username yang unik dan tidak boleh kosong
	Password  string    `gorm:"type:varchar(255);not null"`             // Kolom Password yang tidak boleh kosong
	Phone     string    `gorm:"type:varchar(15);not null"`              // Kolom Phone yang tidak boleh kosong
	Fullname  string    `gorm:"type:varchar(255);not null"`             // Kolom Fullname yang tidak boleh kosong
	CreatedAt time.Time // Kolom CreatedAt
	UpdatedAt time.Time // Kolom UpdatedAt
}

func (User) TableName() string {
	return "custom_schema.users" // Ganti dengan nama schema yang diinginkan
}

type UserRepository interface {
	FindByID(ctx context.Context, id int64) (User, error)
	FindByUsername(ctx context.Context, username string) (User, error)
	InsertUser(ctx context.Context, req dto.UserRegisterReq) (interface{}, error)
}

type UserService interface {
	Authenticate(ctx context.Context, req dto.AuthReq) (dto.BaseResp, error)
	ValidateToken(ctx context.Context, token string) (dto.BaseResp, error)
	Register(ctx context.Context, req dto.UserRegisterReq) (dto.BaseResp, error)
}
