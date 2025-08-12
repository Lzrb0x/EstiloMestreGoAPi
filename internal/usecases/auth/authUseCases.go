package usecases

import (
	"github.com/Lzrb0x/estiloMestreGO/internal/models"
	"github.com/Lzrb0x/estiloMestreGO/internal/repositories"
)

type AuthUseCasesImpl struct {
	RegisterUserUseCase *RegisterUserUseCase
	LoginUserUseCase    *LoginUserUseCase
	RefreshTokenUseCase *RefreshTokenUseCase
}

func NewAuthUseCases(userRepo repositories.UserRepositoryInterface) *AuthUseCasesImpl {
	return &AuthUseCasesImpl{
		RegisterUserUseCase: NewRegisterUserUseCase(userRepo),
		LoginUserUseCase:    NewLoginUserUseCase(userRepo),
		RefreshTokenUseCase: NewRefreshTokenUseCase(userRepo),
	}
}

func (uc *AuthUseCasesImpl) RegisterUser(input RequestRegisterUser) (models.UserResponse, error) {
	return uc.RegisterUserUseCase.RegisterUser(input)
}

func (uc *AuthUseCasesImpl) LoginUser(input RequestLoginUser) (string, string, error) {
	return uc.LoginUserUseCase.LoginUser(input)
}

func (uc *AuthUseCasesImpl) RefreshToken(token string) (string, string, error) {
	return uc.RefreshTokenUseCase.RefreshToken(token)
}
