package dto

import (
	"github.com/google/uuid"
)

type InsertCertificationReq struct {
	UserID         uuid.UUID `json:"user_id"`
	Name           string    `json:"name"`
	Body           string    `json:"body"`
	CredentialID   string    `json:"credential_id"`
	IssueDate      string    `json:"issue_date"`
	ExpirationDate string    `json:"expiration_date"`
}

type UpdateCertificationReq struct {
	Name           string `json:"name"`
	Body           string `json:"body"`
	CredentialID   string `json:"credential_id"`
	IssueDate      string `json:"issue_date"`
	ExpirationDate string `json:"expiration_date"`
}

type CertificationResp struct {
	Certifications []Certification `json:"certifications"`
}

type Certification struct {
	CertificationID uuid.UUID `json:"certification_id"`
	Name            string    `json:"name"`
	Body            string    `json:"body"`
	CredentialID    string    `json:"credential_id"`
	IssueDate       string    `json:"issue_date"`
	ExpirationDate  string    `json:"expiration_date"`
}
