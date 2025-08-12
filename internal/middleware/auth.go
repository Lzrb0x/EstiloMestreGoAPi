package middleware

import (
	"net/http"
	"os"

	"github.com/Lzrb0x/estiloMestreGO/internal/repositories"
	"github.com/Lzrb0x/estiloMestreGO/internal/security"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(userRepo repositories.UserRepositoryInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken, err := c.Cookie("access_token")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		secretKey := os.Getenv("ACCESS_TOKEN_SECRET")
		
		_, claims, err := security.ValidateAccessToken(accessToken, secretKey)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid access token"})
			return
		}

		userIdentifier, ok := claims["userIdentifier"].(string)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token claims"})
			return
		}

		user, err := userRepo.GetByUserIdentifier(userIdentifier)
		if err != nil || user == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		c.Set("userIdentifier", userIdentifier) //inject into the context that will be used by handlers
		c.Next()
	}
}
