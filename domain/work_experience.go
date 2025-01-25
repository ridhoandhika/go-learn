package domain

import (
	"context"
	"ridhoandhika/backend-api/dto"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type WorkExperience struct {
	gorm.Model
	WorkExperienceID uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID           uuid.UUID `gorm:"type:uuid;not null"`
	JobTitle         string    `gorm:"type:varchar(255)"`
	CompanyName      string    `gorm:"type:varchar(255)"`
	StartDate        string    `gorm:"type:varchar(255)"`
	EndDate          string    `gorm:"type:varchar(255)"`
	JobDescription   string    `gorm:"type:text"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func (WorkExperience) TableName() string {
	return "custom_schema.work_experience" // Ganti dengan nama schema yang diinginkan
}

type WorkExperienceRepository interface {
	FindByUserId(ctx context.Context, id uuid.UUID) ([]WorkExperience, error)
	Insert(ctx context.Context, req dto.InsertWorkExperienceReq) (bool, error)
	Update(ctx context.Context, id uuid.UUID, req dto.UpdateWorkExperienceReq) (bool, error)
}

type WorkExperienceService interface {
	FindByUserId(ctx context.Context, id string) (dto.BaseResp, error)
	Insert(ctx context.Context, req dto.InsertWorkExperienceReq) (dto.BaseResp, error)
	Update(ctx context.Context, id string, req dto.UpdateWorkExperienceReq) (dto.BaseResp, error)
}
