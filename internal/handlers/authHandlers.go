package handlers

import (
	"github.com/Lzrb0x/estiloMestreGO/internal/models"
	usecases "github.com/Lzrb0x/estiloMestreGO/internal/usecases/auth"
	"github.com/gin-gonic/gin"
)

type AuthUseCases interface {
	RegisterUser(input usecases.RequestRegisterUser) (models.UserResponse, error)
	LoginUser(input usecases.RequestLoginUser) (string, error)
}

type AuthHandlers struct {
	authUseCases AuthUseCases
}

func NewAuthHandlers(authUC AuthUseCases) *AuthHandlers {
	return &AuthHandlers{
		authUseCases: authUC,
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

	response, err := h.authUseCases.RegisterUser(request)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, response)
}

func (h *AuthHandlers) Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	requestDto := usecases.RequestLoginUser{
		Email:    req.Email,
		Password: req.Password,
	}

	response, err := h.authUseCases.LoginUser(requestDto)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"token": response})

}
