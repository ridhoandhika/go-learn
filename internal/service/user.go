package service

import (
	"context"
	"encoding/json"
	"ridhoandhika/backend-api/domain"
	"ridhoandhika/backend-api/dto"
	"ridhoandhika/backend-api/internal/util"

	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepository  domain.UserRepository
	cacheRepository domain.CacheRepository
}

func NewUser(userRepository domain.UserRepository, cacheRepository domain.CacheRepository) domain.UserService {
	return &userService{
		userRepository:  userRepository,
		cacheRepository: cacheRepository,
	}
}

func (u userService) Authenticate(ctx context.Context, req dto.AuthReq) (dto.AuthResp, error) {
	user, err := u.userRepository.FindByUsername(ctx, req.Username)
	if err != nil {
		return dto.AuthResp{}, err
	}

	if user == (domain.User{}) {
		return dto.AuthResp{}, domain.ErrAuthFailed
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return dto.AuthResp{}, domain.ErrAuthFailed
	}

	token := util.GenerateRandomString(16)

	userJson, _ := json.Marshal(user)
	_ = u.cacheRepository.Set("user:"+token, userJson)

	return dto.AuthResp{
		Token: token,
	}, nil

}

func (u userService) ValidateToken(ctx context.Context, token string) (dto.UserData, error) {
	data, err := u.cacheRepository.Get("user:" + token)
	if err != nil {
		return dto.UserData{}, domain.ErrAuthFailed
	}

	var user domain.User

	_ = json.Unmarshal(data, &user)

	return dto.UserData{
		ID:       user.ID,
		Fullname: user.Fullname,
		Phone:    user.Phone,
		Usename:  user.Usename,
	}, nil
}
