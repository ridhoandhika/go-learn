package dto

import "github.com/google/uuid"

type EducationResp struct {
	Education []Education `json:"education"`
}

type Education struct {
	EducationID  uuid.UUID `json:"education_id"`
	Degree       string    `json:"degree"`
	SchoolName   string    `json:"school_name"`
	FieldOfStudy string    `json:"field_of_study"`
	Description  string    `json:"description"`
	StartDate    string    `json:"start_date"`
	EndDate      string    `json:"end_date"`
}
