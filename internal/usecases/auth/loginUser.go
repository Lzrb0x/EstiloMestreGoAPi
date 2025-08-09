package usecases

import "github.com/Lzrb0x/estiloMestreGO/internal/repositories"

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

func (uc *LoginUserUseCase) LoginUser(request RequestLoginUser) (string, error) {
	return "return token here", nil
}
