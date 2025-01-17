package domain

import (
	"context"
	"ridhoandhika/backend-api/dto"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PersonalInformation struct {
	gorm.Model
	PersonalInfoID uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID         uuid.UUID `gorm:"type:uuid;not null"`
	FirstName      string    `gorm:"type:varchar(255)"`
	LastName       string    `gorm:"type:varchar(255)"`
	PhoneNumber    string    `gorm:"type:varchar(20)"`
	Email          string    `gorm:"type:varchar(100)"`
	Address        string    `gorm:"type:text"`
	Summary        string    `gorm:"type:text"`
	DateOfBirth    string    `gorm:"type:varchar(100)"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (PersonalInformation) TableName() string {
	return "custom_schema.personal_information" // Ganti dengan nama schema yang diinginkan
}

type PersonalInformationRepository interface {
	FindByID(ctx context.Context, id uuid.UUID) (PersonalInformation, error)
	FindByUserID(ctx context.Context, userId uuid.UUID) (PersonalInformation, error)
	Insert(ctx context.Context, req dto.InsertPersonalInformationReq) (interface{}, error)
	Update(ctx context.Context, personalInfoID uuid.UUID, req dto.UpdatePersonalInformationReq) (bool, error)
	// Update(ctx context.Context, req interface{}) (interface{}, error)
}

type PersonalInformationService interface {
	FindByIDPeronalInfo(ctx context.Context, id string) (dto.BaseResp, error)
	Insert(ctx context.Context, req dto.InsertPersonalInformationReq) (dto.BaseResp, error)
	Update(ctx context.Context, personalInfoID uuid.UUID, req dto.UpdatePersonalInformationReq) (dto.BaseResp, error)
	// Update(ctx context.Context, req dto.PersonalInformationReq) (dto.BaseResp, error)
}
