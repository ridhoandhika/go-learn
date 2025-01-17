package domain

import (
	"context"
	"ridhoandhika/backend-api/dto"
	"time"

	"github.com/google/uuid"
)

type WorkExperience struct {
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
	// FindByID(ctx context.Context, id uuid.UUID) (WorkExperience, error)
	// FindByUserID(ctx context.Context, userId uuid.UUID) (WorkExperience, error)
	Insert(ctx context.Context, req dto.InsertWorkExperienceReq) (interface{}, error)
	// Update(ctx context.Context, req interface{}) (interface{}, error)
}

type WorkExperienceService interface {
	// FindByIDPeronalInfo(ctx context.Context, id string) (dto.BaseResp, error)
	Insert(ctx context.Context, req dto.InsertWorkExperienceReq) (dto.BaseResp, error)
	// Update(ctx context.Context, personalInfoID uuid.UUID, req dto.UpdatePersonalInformationReq) (dto.BaseResp, error)
	// Update(ctx context.Context, req dto.PersonalInformationReq) (dto.BaseResp, error)
}
