package dto

import "github.com/google/uuid"

type SkillsResp struct {
	Skills []Skill `json:"skills"`
}
type Skill struct {
	SkillID uuid.UUID `json:"skill_id"`
	Name    string    `json:"name"`
	Level   string    `json:"level"`
}
