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

func User(userRepository domain.UserRepository, cacheRepository domain.CacheRepository) domain.UserService {
	return &userService{
		userRepository:  userRepository,
		cacheRepository: cacheRepository,
	}
}

func (u userService) Authenticate(ctx context.Context, req dto.AuthReq) (dto.BaseResp, error) {
	// get user by username
	user, err := u.userRepository.FindByUsername(ctx, req.Username)
	if err != nil {
		return util.ErrorResponse("404", "User tidak ditemukan", "User not found"), err
	}

	if user == (domain.User{}) {
		return util.ErrorResponse("404", "User tidak ditemukan", "User not found"), err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		// Jika password tidak cocok, kembalikan error 401 Unauthorized
		return util.ErrorResponse("401", "Tidak Berwenang", "Unauthorized"), err
	}

	// generate jwt
	token, err := util.GenerateTokenJWT(req.Username)
	if err != nil {
		return util.ErrorResponse("401", "Tidak Berwenang", "Unauthorized"), err
	}

	userJson, _ := json.Marshal(user)
	_ = u.cacheRepository.Set("user:"+token, userJson)

	return dto.BaseResp{
		ErrorSchema: dto.ErrorSchema{
			ErrorCode: "200",
			ErrorMessage: dto.ErrorMessage{
				Id: "Sukes",
				En: "Success",
			},
		},
		OutputSchema: dto.AuthResp{
			AccessToken: token,
		},
	}, nil

}

func (u userService) ValidateToken(ctx context.Context, token string) (dto.BaseResp, error) {
	data, err := u.cacheRepository.Get("user:" + token)
	if err != nil {
		// return dto.UserData{}, domain.ErrAuthFailed
		return util.ErrorResponse("401", "Tidak Berwenang", "Unauthorized"), err
	}

	var user domain.User

	_ = json.Unmarshal(data, &user)

	return dto.BaseResp{
		ErrorSchema: dto.ErrorSchema{
			ErrorCode: "200",
			ErrorMessage: dto.ErrorMessage{
				Id: "Sukes",
				En: "Success",
			},
		},
		OutputSchema: dto.UserData{
			ID:       user.ID,
			Fullname: user.Fullname,
			Phone:    user.Phone,
			Usename:  user.Usename,
		},
	}, nil
}

func (u userService) Register(ctx context.Context, req dto.UserRegisterReq) (dto.BaseResp, error) {
	// generate hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	if err != nil {
		return util.ErrorResponse("400", "Permintaan Salah", "Bad Request"), err
	}
	// change value to hashed
	req.Password = string(hashedPassword)

	// insert to db
	_, err = u.userRepository.InsertUser(ctx, req)
	if err != nil {
		return util.ErrorResponse("400", "Permintaan Salah", "Bad Request"), err
	}
	// return response
	return util.ErrorResponse("200", "Sukses", "Success"), nil
}
