package domain

import (
	"context"
	"ridhoandhika/backend-api/dto"
)

type User struct {
	ID       int64  `db:"id"`
	Fullname string `db:"fullname"`
	Phone    string `db:"phone"`
	Usename  string `db:"username"`
	Password string `db:"password"`
}

type UserRepository interface {
	FindByID(ctx context.Context, id int64) (User, error)
	FindByUsername(ctx context.Context, username string) (User, error)
	InsertUser(ctx context.Context, req dto.UserRegisterReq) (interface{}, error)
}

type UserService interface {
	Authenticate(ctx context.Context, req dto.AuthReq) (dto.BaseResp, error)
	ValidateToken(ctx context.Context, token string) (dto.BaseResp, error)
	Register(ctx context.Context, req dto.UserRegisterReq) (dto.BaseResp, error)
}
