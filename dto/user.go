package dto

import "github.com/google/uuid"

type UserData struct {
	ID       uuid.UUID `json:"id"`
	Fullname string    `json:"fullname"`
	Phone    string    `json:"phone"`
	Usename  string    `json:"username"`
}

type UserRegisterReq struct {
	Fullname string `json:"fullname"`
	Phone    string `json:"phone"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRegisterResp struct {
	Message string `json:"message"`
}

type UserEducationResp struct {
	Education []Education `json:"educations"`
}

type UserWorkExperienceResp struct {
	WorkExperience []WorkExperience `json:"work_experiences"`
}

type UserSkillResp struct {
	Skill []Skill `json:"skills"`
}

type UserPersonalInformationResp struct {
	PersonalInformation PersonalInformationResp `json:"personal_information"`
}

type UserCertificationResp struct {
	Certification []Certification `json:"certifications"`
}
