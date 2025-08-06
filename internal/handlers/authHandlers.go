package handlers

import (
	"github.com/Lzrb0x/estiloMestreGO/internal/models"
	usecases "github.com/Lzrb0x/estiloMestreGO/internal/useCases/auth"
	"github.com/gin-gonic/gin"
)

type AuthUseCases interface {
	Execute(input usecases.RequestRegisterUser) (models.UserResponse, error)
}

type AuthHandlers struct {
	registerUserUseCase AuthUseCases
}

func NewAuthHandlers(registerUC AuthUseCases) *AuthHandlers{
	return &AuthHandlers{
		registerUserUseCase: registerUC,
	}
}

func (h *AuthHandlers) Register(c *gin.Context) {
	var req struct {
		Name  string `json:"name" binding:"required"`
		Email string `json:"email" binding:"required,email"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	request := usecases.RequestRegisterUser{
		Name:  req.Name,
		Email: req.Email,
	}

	response, err := h.registerUserUseCase.Execute(request)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}


	c.JSON(201, response)
}


