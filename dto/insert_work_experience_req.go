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
