package domain

import (
	"context"
	"ridhoandhika/backend-api/dto"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID             uuid.UUID           `gorm:"type:uuid;primaryKey"`       // UUID sebagai primary key
	Username       string              `gorm:"type:varchar(100);not null"` // Kolom Username yang unik dan tidak boleh kosong
	Password       string              `gorm:"type:varchar(255);not null"` // Kolom Password yang tidak boleh kosong
	Phone          string              `gorm:"type:varchar(15);not null"`  // Kolom Phone yang tidak boleh kosong
	Fullname       string              `gorm:"type:varchar(255);not null"` // Kolom Fullname yang tidak boleh kosong
	PersonalInfo   PersonalInformation `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	WorkExperience []WorkExperience    `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Education      []Education         `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Skill          []Skill             `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Certification  []Certification     `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	CreatedAt      time.Time           // Kolom CreatedAt
	UpdatedAt      time.Time           // Kolom UpdatedAt
}

func (User) TableName() string {
	return "custom_schema.users" // Ganti dengan nama schema yang diinginkan
}

type UserRepository interface {
	FindByID(ctx context.Context, id uuid.UUID) (User, error)
	FindByUsername(ctx context.Context, username string) (User, error)
	InsertUser(ctx context.Context, req dto.UserRegisterReq) (interface{}, error)
	UserWithEducation(ctx context.Context, userID uuid.UUID) (*User, error)
	UserWithWorkExperience(ctx context.Context, userID uuid.UUID) (*User, error)
	UserWithSkill(ctx context.Context, userID uuid.UUID) (*User, error)
	UserWithPersonalInformation(ctx context.Context, userID uuid.UUID) (*User, error)
	UserWithCertification(ctx context.Context, userID uuid.UUID) (*User, error)
}

type UserService interface {
	GetUser(ctx context.Context, userID string) (dto.BaseResp, error)
	GetEducation(ctx context.Context, userID string) (dto.BaseResp, error)
	GetPersonalInformation(ctx context.Context, userID string) (dto.BaseResp, error)
	GetWorkExperience(ctx context.Context, userID string) (dto.BaseResp, error)
	GetSkill(ctx context.Context, userID string) (dto.BaseResp, error)
	GetCertification(ctx context.Context, userID string) (dto.BaseResp, error)
}
