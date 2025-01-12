package dto

import "github.com/google/uuid"

type UserData struct {
	ID       uuid.UUID `json:"id"`
	Fullname string    `json:"fullname"`
	Phone    string    `json:"phone"`
	Usename  string    `json:"username"`
}
