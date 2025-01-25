package domain

import (
	"context"
	"ridhoandhika/backend-api/dto"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Level string

const (
	Beginner     Level = "Beginner"
	Intermediate Level = "Intermediate"
	Advanced     Level = "Advanced"
)

type Skill struct {
	gorm.Model
	SkillID   uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID    uuid.UUID `gorm:"type:uuid;not null"`
	Name      string    `gorm:"type:varchar(100);not null"`
	Level     Level     `gorm:"type:level"` // GORM enum definition
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Skill) TableName() string {
	return "custom_schema.skill" // Ganti dengan nama schema yang diinginkan
}

type SkillRepository interface {
	FindByUserId(ctx context.Context, id uuid.UUID) ([]Skill, error)
	Insert(ctx context.Context, req dto.InsertSkillReq) (bool, error)
	Update(ctx context.Context, id uuid.UUID, req dto.UpdateSkillReq) (bool, error)
}

type SkillService interface {
	FindByUserId(ctx context.Context, id string) (dto.BaseResp, error)
	Insert(ctx context.Context, req dto.InsertSkillReq) (dto.BaseResp, error)
	Update(ctx context.Context, id string, req dto.UpdateSkillReq) (dto.BaseResp, error)
}
