package usecases

import (
	"errors"

	"github.com/Lzrb0x/estiloMestreGO/internal/models"
	"github.com/Lzrb0x/estiloMestreGO/internal/repositories"
	"github.com/Lzrb0x/estiloMestreGO/internal/security"
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
	Name     string
	Email    string
	Password string
}

func (uc *RegisterUserUseCase) RegisterUser(request RequestRegisterUser) (models.UserResponse, error) {

	user, err := uc.userRepo.GetUserByEmail(request.Email)
	if err != nil {
		return models.UserResponse{}, err
	}

	if user != nil {
		return models.UserResponse{}, errors.New("user already exists")
	}

	//hash the password

	hashedPassword, err := security.HashPassword(request.Password)
	if err != nil {
		return models.UserResponse{}, errors.New("failed to hash password")
	}

	newUser, err := models.NewUser(request.Name, request.Email, hashedPassword)
	if err != nil {
		return models.UserResponse{}, err
	}

	err = uc.userRepo.AddUser(newUser)
	if err != nil {
		return models.UserResponse{}, errors.New("failed to register user")
	}

	return newUser.ToResponse(), nil
}
