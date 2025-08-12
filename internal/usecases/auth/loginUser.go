package usecases

import (
	"errors"

	"github.com/Lzrb0x/estiloMestreGO/internal/repositories"
	"github.com/Lzrb0x/estiloMestreGO/internal/security"
)

type LoginUserUseCase struct {
	userRepo repositories.UserRepositoryInterface
}

func NewLoginUserUseCase(userRepo repositories.UserRepositoryInterface) *LoginUserUseCase {
	return &LoginUserUseCase{
		userRepo: userRepo,
	}
}

type RequestLoginUser struct {
	Email    string
	Password string
}

func (uc *LoginUserUseCase) LoginUser(request RequestLoginUser) (string, string, error) {

	var user, err = uc.userRepo.GetUserByEmail(request.Email)
	if err != nil {
		return "", "", err
	}

	if user == nil {
		return "", "", errors.New("email or password invalid")
	}

	if !security.ValidatePassword(user.Password, request.Password) {
		return "", "", errors.New("email or password invalid")
	}

	accessToken, err := security.GenerateAccessToken(user.UserIdentifier)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := security.GenerateRefreshToken(user.UserIdentifier)
	if err != nil {
		return "", "", err
	}

	user.AddRefreshToken(refreshToken)
	err = uc.userRepo.UpdateUser(user)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
