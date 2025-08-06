package usecases

import (
	"github.com/Lzrb0x/estiloMestreGO/internal/models"
	"github.com/Lzrb0x/estiloMestreGO/internal/repositories"
)

type RegisterUserUseCase struct {
	userRepo repositories.UserRepositoryInterface
}

func NewRegisterUserUseCase(userRepo repositories.UserRepositoryInterface) *RegisterUserUseCase {
	return &RegisterUserUseCase{
		userRepo: userRepo,
	}
}

type RequestRegisterUser struct {
	Name  string
	Email string
}

func (uc *RegisterUserUseCase) Execute(request RequestRegisterUser) (models.UserResponse, error) {

	user, err := models.NewUser(request.Name, request.Email)
	if err != nil {
		return models.UserResponse{}, err
	}

	err = uc.userRepo.AddUser(user)
	if err != nil {
		return models.UserResponse{}, err
	}

	return user.ToResponse(), nil
}
