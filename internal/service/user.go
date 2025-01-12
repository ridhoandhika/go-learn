package service

import (
	"context"
	"errors"
	"ridhoandhika/backend-api/domain"
	"ridhoandhika/backend-api/dto"
	"ridhoandhika/backend-api/internal/util"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepository domain.UserRepository
}

func User(userRepository domain.UserRepository) domain.UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (u userService) Authenticate(ctx context.Context, req dto.AuthReq) (dto.AuthResp, error) {
	// get user by username
	user, err := u.userRepository.FindByUsername(ctx, req.Username)
	if err != nil {
		return dto.AuthResp{}, errors.New("401")
		// return util.ErrorResponse("404", "User tidak ditemukan", "User not found"), nil
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		// Jika password tidak cocok, kembalikan error 401 Unauthorized
		return dto.AuthResp{}, errors.New("401")
	}

	// generate jwt
	token, err := util.GenerateTokenJWT(req.Username)
	if err != nil {
		return dto.AuthResp{}, errors.New("401")
	}

	return dto.AuthResp{
		AccessToken: token,
	}, nil

}

func (u userService) ValidateToken(ctx context.Context, token string) (dto.AuthResp, error) {
	// Memverifikasi token menggunakan VerifyToken
	tokenResp, err := util.VerifyToken(token)
	if err != nil {
		// Jika token tidak valid, mengembalikan response error
		return dto.AuthResp{}, errors.New("invalid")
	}

	claims, ok := tokenResp.Claims.(jwt.MapClaims)
	if !ok {
		return dto.AuthResp{}, errors.New("invalid")
	}

	// Mengambil nilai 'exp' dari klaim
	username, ok := claims["username"].(string) // Klaim 'exp' biasanya bertipe float64
	if !ok {
		return dto.AuthResp{}, errors.New("invalid")
	}

	newToken, err := util.GenerateTokenJWT(username)
	if err != nil {
		return dto.AuthResp{}, errors.New("invalid")
	}

	// Mengembalikan response dengan data yang diinginkan
	return dto.AuthResp{
		AccessToken: newToken,
	}, nil
}

func (u userService) Register(ctx context.Context, req dto.UserRegisterReq) (dto.BaseResp, error) {
	// generate hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	if err != nil {
		return util.ErrorResponse("400", "Permintaan Tidak Valid", "Bad Request"), nil
	}
	// change value to hashed
	req.Password = string(hashedPassword)

	// insert to db
	_, err = u.userRepository.InsertUser(ctx, req)
	if err != nil {
		return util.ErrorResponse("400", "Permintaan Tidak Valid", "Bad Request"), nil
	}
	// return response
	return util.ErrorResponse("200", "Sukses", "Success"), nil
}
