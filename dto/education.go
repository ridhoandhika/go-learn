package dto

import "github.com/google/uuid"

type InsertEducationReq struct {
	EducationID  uuid.UUID `json:"education_id"`
	UserID       uuid.UUID `json:"user_id"`
	Degree       string    `json:"degree"`
	SchoolName   string    `json:"school_name"`
	FieldOfStudy string    `json:"field_of_study"`
	Description  string    `json:"description"`
	StartDate    string    `json:"start_date"`
	EndDate      string    `json:"end_date"`
}

type UpdateEducationReq struct {
	Degree       string `json:"degree"`
	SchoolName   string `json:"school_name"`
	FieldOfStudy string `json:"field_of_study"`
	Description  string `json:"description"`
	StartDate    string `json:"start_date"`
	EndDate      string `json:"end_date"`
}

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
