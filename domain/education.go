package domain

import (
	"context"
	"ridhoandhika/backend-api/dto"
	"time"

	"github.com/google/uuid"
)

type Education struct {
	EducationID  uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID       uuid.UUID `gorm:"type:uuid;not null"`
	Degree       string    `gorm:"type:varchar(255)"`
	SchoolName   string    `gorm:"type:varchar(255)"`
	FieldOfStudy string    `gorm:"type:varchar(255)"`
	StartDate    string    `gorm:"type:varchar(255)"`
	EndDate      string    `gorm:"type:varchar(255)"`
	Description  string    `gorm:"type:text"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (Education) TableName() string {
	return "custom_schema.education" // Ganti dengan nama schema yang diinginkan
}

type EducationRepository interface {
	FindByUserId(ctx context.Context, id uuid.UUID) ([]Education, error)
	Insert(ctx context.Context, req dto.InsertEducationReq) (bool, error)
	Update(ctx context.Context, id uuid.UUID, req dto.UpdateEducationReq) (bool, error)
}

type EducationService interface {
	FindByUserId(ctx context.Context, id string) (dto.BaseResp, error)
	Insert(ctx context.Context, req dto.InsertEducationReq) (dto.BaseResp, error)
	Update(ctx context.Context, id string, req dto.UpdateEducationReq) (dto.BaseResp, error)
}
