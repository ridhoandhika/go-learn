package service

import (
	"context"
	"ridhoandhika/backend-api/domain"
	"ridhoandhika/backend-api/dto"
	"ridhoandhika/backend-api/internal/util"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepository  domain.UserRepository
	cacheRepository domain.CacheRepository
}

func User(userRepository domain.UserRepository) domain.UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (u userService) Authenticate(ctx context.Context, req dto.AuthReq) (dto.BaseResp, error) {
	// get user by username
	user, err := u.userRepository.FindByUsername(ctx, req.Username)
	if err != nil {
		return util.ErrorResponse("404", "User tidak ditemukan", "User not found"), nil
	}

	if user == (domain.User{}) {
		return util.ErrorResponse("404", "User tidak ditemukan", "User not found"), err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		// Jika password tidak cocok, kembalikan error 401 Unauthorized
		return util.ErrorResponse("401", "Tidak Berwenang", "Unauthorized"), nil
	}

	// generate jwt
	token, err := util.GenerateTokenJWT(req.Username)
	if err != nil {
		return util.ErrorResponse("401", "Tidak Berwenang", "Unauthorized"), err
	}

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
	// Memverifikasi token menggunakan VerifyToken
	tokenResp, err := util.VerifyToken(token)
	if err != nil {
		// Jika token tidak valid, mengembalikan response error
		return util.ErrorResponse("401", "Tidak Berwenang", "Unauthorized"), nil
	}

	claims, ok := tokenResp.Claims.(jwt.MapClaims)
	if !ok {
		return util.ErrorResponse("500", "Internal Server Error", "Unable to parse claims"), nil
	}

	// Mengambil nilai 'exp' dari klaim
	username, ok := claims["username"].(string) // Klaim 'exp' biasanya bertipe float64
	if !ok {
		return util.ErrorResponse("500", "Internal Server Error", "Expiration time not found"), nil
	}

	newToken, err := util.GenerateTokenJWT(username)
	if err != nil {
		return util.ErrorResponse("401", "Tidak Berwenang", "Unauthorized"), nil
	}

	// Mengembalikan response dengan data yang diinginkan
	return dto.BaseResp{
		ErrorSchema: dto.ErrorSchema{
			ErrorCode: "200",
			ErrorMessage: dto.ErrorMessage{
				Id: "Sukses", // Perbaiki pengejaan dari "Sukes" menjadi "Sukses"
				En: "Success",
			},
		},
		OutputSchema: dto.AuthResp{
			AccessToken: newToken,
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
