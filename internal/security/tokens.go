package security

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func GenerateAccessToken(userIdentifier uuid.UUID) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["userIdentifier"] = userIdentifier.String()
	claims["exp"] = time.Now().Add(15 * time.Minute).Unix()

	secret := os.Getenv("ACCESS_TOKEN_SECRET")
	if secret == "" {
		return "", errors.New("ACCESS_TOKEN_SECRET is not set")
	}
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", errors.New("failed to generate access token")
	}

	return tokenString, nil
}

func GenerateRefreshToken(userIdentifier uuid.UUID) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["userIdentifier"] = userIdentifier.String()
	claims["exp"] = time.Now().Add(7 * 24 * time.Hour).Unix()

	secret := os.Getenv("REFRESH_TOKEN_SECRET")
	if secret == "" {
		return "", errors.New("REFRESH_TOKEN_SECRET is not set")
	}
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", errors.New("failed to generate refresh token")
	}

	return tokenString, nil
}

func ValidateAccessToken(tokenString, secretKey string) (*jwt.Token, jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, nil, err
	}
	
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, nil, errors.New("invalid token")
	}

	return token, claims, nil
}
