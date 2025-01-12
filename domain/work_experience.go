package domain

import (
	"time"

	"github.com/google/uuid"
)

type WorkExperience struct {
	WorkExperienceID uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID           uuid.UUID `gorm:"type:uuid;not null"`
	JobTitle         string    `gorm:"type:varchar(255)"`
	CompanyName      string    `gorm:"type:varchar(255)"`
	StartDate        time.Time
	EndDate          *time.Time
	JobDescription   string `gorm:"type:text"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func (WorkExperience) TableName() string {
	return "custom_schema.work_experience" // Ganti dengan nama schema yang diinginkan
}
