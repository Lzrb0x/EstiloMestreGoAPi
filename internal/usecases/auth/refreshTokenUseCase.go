package usecases

import (
	"github.com/Lzrb0x/estiloMestreGO/internal/repositories"
	"github.com/Lzrb0x/estiloMestreGO/internal/security"
)

type RefreshTokenUseCase struct {
	userRepo repositories.UserRepositoryInterface
}

func NewRefreshTokenUseCase(userRepo repositories.UserRepositoryInterface) *RefreshTokenUseCase {
	return &RefreshTokenUseCase{
		userRepo: userRepo,
	}
}

func (u *RefreshTokenUseCase) RefreshToken(token string) (string, string, error) {

	user, err := u.userRepo.GetUserByRefreshToken(token)
	if err != nil {
		return "", "", err
	}

	if user == nil {
		return "", "", nil
	}

	newAccessToken, err := security.GenerateAccessToken(user.UserIdentifier)
	if err != nil {
		return "", "", err
	}

	newRefreshToken, err := security.GenerateRefreshToken(user.UserIdentifier)
	if err != nil {
		return "", "", err
	}

	user.AddRefreshToken(newRefreshToken)
	err = u.userRepo.UpdateUser(user)
	if err != nil {
		return "", "", err
	}

	return newAccessToken, newRefreshToken, nil
}
