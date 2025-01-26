package dto

import (
	"github.com/google/uuid"
)

type InsertSkillReq struct {
	UserID uuid.UUID `json:"user_id"`
	Name   string    `json:"name"`
	Level  string    `json:"level"`
}

type UpdateSkillReq struct {
	Name  string `json:"name"`
	Level string `json:"level"`
}

type SkillsResp struct {
	Skills []Skill `json:"skills"`
}

type Skill struct {
	SkillID uuid.UUID `json:"skill_id"`
	Name    string    `json:"name"`
	Level   string    `json:"level"`
}
