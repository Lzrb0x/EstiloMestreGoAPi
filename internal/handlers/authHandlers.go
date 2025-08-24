package handlers

import (
	"net/http"

	"github.com/Lzrb0x/estiloMestreGO/internal/models"
	usecases "github.com/Lzrb0x/estiloMestreGO/internal/usecases/auth"
	"github.com/gin-gonic/gin"
)

type AuthUseCases interface {
	RegisterUser(input usecases.RequestRegisterUser) (models.UserResponse, error)
	LoginUser(input usecases.RequestLoginUser) (string, string, error)
	RefreshToken(token string) (string, string, error)
}

type AuthHandlers struct {
	authUseCases AuthUseCases
}

func NewAuthHandlers(authUC AuthUseCases) *AuthHandlers {
	return &AuthHandlers{
		authUseCases: authUC,
	}
}

// Register godoc
// @Summary Register a new user
// @Description Register a new user with name, email, and password
// @Tags auth
// @Accept  json
// @Produce  json
// @Param   user     body      object{name=string,email=string,password=string}   true  "User Credentials"
// @Success 201 {object} models.UserResponse
// @Failure 400 {object} object{error=string}
// @Router /auth/register [post]
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

	requestDto := usecases.RequestRegisterUser{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	response, err := h.authUseCases.RegisterUser(requestDto)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, response)
}

// Login godoc
// @Summary Login a user
// @Description Login a user with email and password
// @Tags auth
// @Accept  json
// @Produce  json
// @Param   user     body      object{email=string,password=string}   true  "User Credentials"
// @Success 200 {object} object{token=string}
// @Failure 400 {object} object{error=string}
// @Router /auth/login [post]
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

	accessToken, refreshToken, err := h.authUseCases.LoginUser(requestDto)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie(
		"access_token",
		accessToken,
		15*60,
		"/",
		"localhost",
		false, //for development purposes
		true,
	)

	c.SetCookie(
		"refresh_token",
		refreshToken,
		7*24*60*60,
		"/",
		"localhost",
		false, //for development purposes
		true,
	)

	c.JSON(http.StatusOK, gin.H{"success": true})
}

func (h *AuthHandlers) Refresh(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	newAccessToken, newRefreshToken, err := h.authUseCases.RefreshToken(refreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	c.SetCookie(
		"access_token",
		newAccessToken,
		15*60,
		"/",
		"localhost",
		false, //for development purposes
		true,
	)

	c.SetCookie(
		"refresh_token",
		newRefreshToken,
		7*24*60*60,
		"/",
		"localhost",
		false, //for development purposes
		true,
	)

	c.JSON(http.StatusOK, gin.H{"success": true})
}

func (h *AuthHandlers) Logout(c *gin.Context) {
	c.SetCookie("access_token", "", -1, "/", "localhost", false, true)
	c.SetCookie("refresh_token", "", -1, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{"success": true})
}
