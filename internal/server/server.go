package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Lzrb0x/estiloMestreGO/internal/handlers"
	"github.com/Lzrb0x/estiloMestreGO/internal/repositories"
	usecases "github.com/Lzrb0x/estiloMestreGO/internal/usecases/auth"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func NewServer(db *gorm.DB) *http.Server {

	port := "8080" //get from env or config

	// repositories
	userRepository := repositories.NewUserRepository(db)


	// use cases
	authUseCases := usecases.NewAuthUseCases(userRepository)


	return &http.Server{
		Addr: fmt.Sprintf(":%s", port),
		Handler: RegisterRoutes(authUseCases),
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}

func RegisterRoutes(authUC *usecases.AuthUseCasesImpl) http.Handler {
	r := gin.Default()

	authHandlers := handlers.NewAuthHandlers(authUC)

	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/register", authHandlers.Register)
		authRoutes.POST("/login", authHandlers.Login)
	}

	return r
}
