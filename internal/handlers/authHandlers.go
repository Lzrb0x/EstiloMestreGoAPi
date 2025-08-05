package handlers

import (
	"github.com/Lzrb0x/estiloMestreGO/internal/repositories"
	useCases "github.com/Lzrb0x/estiloMestreGO/internal/useCases/auth/register"
	"github.com/gin-gonic/gin"
)

type AuthHandlers struct {
	userRepo repositories.UserRepositoryInterface
}

func NewAuthHandlers(userRepo repositories.UserRepositoryInterface) *AuthHandlers {
	return &AuthHandlers{
		userRepo: userRepo,
	}
}

func (h *AuthHandlers) Register(c *gin.Context) {
	user, err := useCases.NewRegisterUseCase(h.userRepo).Execute()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, user)
}

