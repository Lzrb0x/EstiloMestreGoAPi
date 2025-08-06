package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Lzrb0x/estiloMestreGO/internal/handlers"
	"github.com/Lzrb0x/estiloMestreGO/internal/repositories"
	usecases "github.com/Lzrb0x/estiloMestreGO/internal/useCases/auth"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func NewServer(db *gorm.DB) *http.Server {
	port := "8080"

	userRepository := repositories.NewUserRepository(db)

	registerUserUseCase := usecases.NewRegisterUserUseCase(userRepository)


	return &http.Server{
		Addr: fmt.Sprintf(":%s", port),
		Handler: RegisterRoutes(registerUserUseCase),
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}

func RegisterRoutes(registerUC *usecases.RegisterUserUseCase) http.Handler {
	r := gin.Default()

	authHandlers := handlers.NewAuthHandlers(registerUC)

	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/register", authHandlers.Register)
	}

	return r
}
