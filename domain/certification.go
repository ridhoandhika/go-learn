package domain

import (
	"context"
	"ridhoandhika/backend-api/dto"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Certification struct {
	gorm.Model
	CertificationID uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID          uuid.UUID `gorm:"type:uuid;not null"`
	Name            string    `gorm:"type:varchar(255)"`
	Body            string    `gorm:"type:text"`
	CredentialID    string    `gorm:"type:varchar(255)"`
	IssueDate       *time.Time
	ExpirationDate  *time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func (Certification) TableName() string {
	return "custom_schema.certification" // Ganti dengan nama schema yang diinginkan
}

type CertificationRepository interface {
	FindByUserId(ctx context.Context, id uuid.UUID) ([]Certification, error)
	Insert(ctx context.Context, req dto.InsertCertificationReq) (bool, error)
	Update(ctx context.Context, id uuid.UUID, req dto.UpdateCertificationReq) (bool, error)
}

type CertificationService interface {
	FindByUserId(ctx context.Context, id string) (dto.BaseResp, error)
	Insert(ctx context.Context, req dto.InsertCertificationReq) (dto.BaseResp, error)
	Update(ctx context.Context, id string, req dto.UpdateCertificationReq) (dto.BaseResp, error)
}
