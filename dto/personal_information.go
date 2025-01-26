package dto

import (
	"github.com/google/uuid"
)

type InsertPersonalInformationReq struct {
	UserID      uuid.UUID `json:"user_id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email"`
	Address     string    `json:"address"`
	Summary     string    `json:"summary"`
	DateOfBirth string    `json:"date_of_birth"`
}

type UpdatePersonalInformationReq struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	Summary     string `json:"summary"`
	DateOfBirth string `json:"date_of_birth"`
}

type PersonalInformationResp struct {
	PersonalInfoID uuid.UUID `json:"personal_info_id"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	PhoneNumber    string    `json:"phone_number"`
	Email          string    `json:"email"`
	Address        string    `json:"address"`
	Summary        string    `json:"summary"`
	DateOfBirth    string    `json:"date_of_birth"`
}
