package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model               //id, created_at, updated_at, deleted_at
	UserIdentifier uuid.UUID `gorm:"type:uuid" json:"user_identifier"`
	RefreshToken   string    `gorm:"type:varchar(512)" json:"refresh_token,omitempty"`
	Name           string    `gorm:"not null;type:varchar(255)" json:"name"`
	Email          string    `gorm:"type:varchar(255);unique" json:"email"`
	Password       string    `gorm:"type:varchar(512)" json:"-"` // "-" omite do JSON por segurança
}

func NewUser(name, email, password string) (*User, error) {
	user := &User{
		Name:           name,
		Email:          email,
		Password:       password,
		UserIdentifier: uuid.New(),
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) AddRefreshToken(token string) {
	u.RefreshToken = token
}

func (u *User) Validate() error {

	var errorMessages []string

	if u.Name == "" {
		errorMessages = append(errorMessages, "Name is required")
	}
	if u.Email == "" {
		errorMessages = append(errorMessages, "Email is required")
	}
	if len(errorMessages) > 0 {
		return fmt.Errorf("%s", strings.Join(errorMessages, ", "))
	}

	return nil
}

type UserResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (u *User) ToResponse() UserResponse {
	return UserResponse{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt.Format(time.RFC3339),
		UpdatedAt: u.UpdatedAt.Format(time.RFC3339),
	}
}
