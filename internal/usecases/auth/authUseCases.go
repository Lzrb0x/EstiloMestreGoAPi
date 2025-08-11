package usecases

import (
	"github.com/Lzrb0x/estiloMestreGO/internal/models"
	"github.com/Lzrb0x/estiloMestreGO/internal/repositories"
)

type AuthUseCasesImpl struct {
	RegisterUserUseCase *RegisterUserUseCase
	LoginUserUseCase    *LoginUserUseCase
}

func NewAuthUseCases(userRepo repositories.UserRepositoryInterface) *AuthUseCasesImpl {
	return &AuthUseCasesImpl{
		RegisterUserUseCase: NewRegisterUserUseCase(userRepo),
		LoginUserUseCase:    NewLoginUserUseCase(userRepo),
	}
}

func (uc *AuthUseCasesImpl) RegisterUser(input RequestRegisterUser) (models.UserResponse, error) {
	return uc.RegisterUserUseCase.RegisterUser(input)
}

func (uc *AuthUseCasesImpl) LoginUser(input RequestLoginUser) (string, error) {
	return uc.LoginUserUseCase.LoginUser(input)
}
