package dto

type UserRegisterReq struct {
	Fullname string `json:"fullname"`
	Phone    string `json:"phone"`
	Username string `json:"username"`
	Password string `json:"password"`
}
