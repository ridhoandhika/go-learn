package dto

type UpdateEducationReq struct {
	Degree       string `json:"degree"`
	SchoolName   string `json:"school_name"`
	FieldOfStudy string `json:"field_of_study"`
	Description  string `json:"description"`
	StartDate    string `json:"start_date"`
	EndDate      string `json:"end_date"`
}
