package useCases

import (
	"github.com/Lzrb0x/estiloMestreGO/internal/models"
	"github.com/Lzrb0x/estiloMestreGO/internal/repositories"
)

type RegisterUseCase struct {
	userRepo repositories.UserRepositoryInterface
}

func NewRegisterUseCase(userRepo repositories.UserRepositoryInterface) *RegisterUseCase {
	return &RegisterUseCase{
		userRepo: userRepo,
	}
}

func (uc *RegisterUseCase) Execute() (*models.User, error) {
	user := models.User{
		Name:  "exampleUser",
		Email: "example@example.com",
	}
	err := uc.userRepo.AddUser(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
