package handlers

import (
	"github.com/Lzrb0x/estiloMestreGO/internal/models"
	"github.com/Lzrb0x/estiloMestreGO/internal/repositories"
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

	var req struct {
		Name  string `json:"name" binding:"required"`
		Email string `json:"email" binding:"required,email"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request data"})
		return
	}

	_, err := models.NewUser(req.Name, req.Email)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User registered successfully"})
}

func (h *AuthHandlers) Login(c *gin.Context) {
	var req struct {
		Name  string `json:"name" binding:"required"`
		Email string `json:"email" binding:"required,email"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request data"})
		return
	}

}
