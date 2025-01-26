package dto

import (
	"github.com/google/uuid"
)

type InsertWorkExperienceReq struct {
	UserID         uuid.UUID `json:"user_id"`
	JobTitle       string    `json:"job_title"`
	CompanyName    string    `json:"company_name"`
	StartDate      string    `json:"start_date"`
	EndDate        string    `json:"end_date"`
	JobDescription string    `json:"job_description"`
}

type UpdateWorkExperienceReq struct {
	JobTitle       string `json:"job_title"`
	CompanyName    string `json:"company_name"`
	StartDate      string `json:"start_date"`
	EndDate        string `json:"end_date"`
	JobDescription string `json:"job_description"`
}

type WorkExperiencesResp struct {
	WorkExperience []WorkExperience `json:"work_experience"`
}
type WorkExperience struct {
	WorkExperienceID uuid.UUID `json:"work_experience_id"`
	JobTitle         string    `json:"job_title"`
	CompanyName      string    `json:"company_name"`
	StartDate        string    `json:"start_date"`
	EndDate          string    `json:"end_date"`
	JobDescription   string    `json:"job_description"`
}
