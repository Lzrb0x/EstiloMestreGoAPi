package handlers

import (
	"github.com/Lzrb0x/estiloMestreGO/internal/models"
	usecases "github.com/Lzrb0x/estiloMestreGO/internal/usecases/auth"
	"github.com/gin-gonic/gin"
)

type AuthUseCases interface {
	RegisterUser(input usecases.RequestRegisterUser) (models.UserResponse, error)
}

type AuthHandlers struct {
	registerUserUseCase AuthUseCases
}

func NewAuthHandlers(authUC AuthUseCases) *AuthHandlers {
	return &AuthHandlers{
		registerUserUseCase: authUC,
	}
}

func (h *AuthHandlers) Register(c *gin.Context) {
	var req struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	request := usecases.RequestRegisterUser{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	response, err := h.registerUserUseCase.RegisterUser(request)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, response)
}
