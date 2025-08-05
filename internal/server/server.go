package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Lzrb0x/estiloMestreGO/internal/handlers"
	"github.com/Lzrb0x/estiloMestreGO/internal/repositories"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	userRepo repositories.UserRepositoryInterface
}

func NewServer(db *gorm.DB) *http.Server {
	port := "8080"

	userRepository := repositories.NewUserRepository(db)

	server := &Server{
		userRepo: userRepository,
	}

	return &http.Server{
		Addr: fmt.Sprintf(":%s", port),
		Handler: server.RegisterRoutes(),
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	authHandlers := handlers.NewAuthHandlers(s.userRepo)

	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/register", authHandlers.Register)
	}

	return r
}
