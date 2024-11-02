package dto

type UserData struct {
	ID       int64  `json:"id"`
	Fullname string `json:"fullname"`
	Phone    string `json:"phone"`
	Usename  string `json:"username"`
}
