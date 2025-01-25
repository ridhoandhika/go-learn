package dto

type WorkExperiencesResp struct {
	WorkExperience []WorkExperience `json:"work_experience"`
}
type WorkExperience struct {
	JobTitle       string `json:"job_title"`
	CompanyName    string `json:"company_name"`
	StartDate      string `json:"start_date"`
	EndDate        string `json:"end_date"`
	JobDescription string `json:"job_description"`
}
